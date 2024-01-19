package rocketmq

import (
	"context"
	"fmt"

	v2 "github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-bamboo/pkg/log"
	otelext "github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type (
	ConsumeHandle func(ctx context.Context, topic string, key, message []byte) error

	ConsumeHandler interface {
		Consume(ctx context.Context, topic string, key, message []byte) error
	}
	rocketQueue struct {
		c       *Conf
		handler ConsumeHandler

		sub        v2.PushConsumer
		tracer     trace.Tracer
		propagator propagation.TextMapPropagator
		subCounter metrics.Counter // 发送次数

		//wg  sync.WaitGroup
		//ctx context.Context
		//cf  context.CancelFunc
	}

	rocketQueues struct {
		queues []*rocketQueue
	}
)

func MustNewQueue(c *Conf, handler ConsumeHandler) queue.MessageQueue {
	q, err := NewQueue(c, handler)
	if err != nil {
		log.Fatal(err)
	}
	return q
}

func NewQueue(c *Conf, handler ConsumeHandler) (queue.MessageQueue, error) {
	q := rocketQueues{}
	cc, err := newKafkaQueue(c, handler)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	q.queues = append(q.queues, cc)
	return &q, nil
}

func (q rocketQueues) Name() string {
	return "rocketmq"
}

func (q rocketQueues) Start(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Start(ctx)
	}
	return nil
}

func (q rocketQueues) Stop(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Stop(ctx)
	}
	return nil
}

func newKafkaQueue(config *Conf, handler ConsumeHandler) (k *rocketQueue, err error) {
	model := consumer.Clustering
	if config.Broadcast {
		model = consumer.BroadCasting
	}
	namesrvAdd, err := primitive.NewNamesrvAddr(config.Addrs...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	cs, err := v2.NewPushConsumer(
		consumer.WithGroupName(config.GroupId),
		consumer.WithNameServer(namesrvAdd),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: config.AccessKey,
			SecretKey: config.SecretKey,
		}),
		consumer.WithConsumerModel(model),
		//consumer.WithNamespace(config.Namespace),
		consumer.WithConsumeMessageBatchMaxSize(1), //
	)
	if err != nil {
		return nil, errors.FromError(err)
	}
	//ctx, cf := context.WithCancel(context.Background())
	k = &rocketQueue{
		c:       config,
		handler: handler,

		sub:        cs,
		tracer:     otel.Tracer("rocketmq"),
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),

		//ctx: ctx,
		//cf:  cf,
	}
	return
}

func (c *rocketQueue) Start(context.Context) error {
	log.Infof("start cunsumer topic:%v", c.c.Topic)
	c.consumGroupTopic(c.c.Topic, c.c.Expression)
	if err := c.sub.Start(); err != nil {
		log.Error(err)
		return err
	}
	log.Infof("start rocket consumer.")
	return nil
}

func (c *rocketQueue) Stop(context.Context) error {
	if err := c.sub.Shutdown(); err != nil {
		return err
	}
	log.Infof("stop rocket consumer.")
	return nil
}

func (c *rocketQueue) consumGroupTopic(topic, expression string) {
	selector := consumer.MessageSelector{Type: consumer.TAG, Expression: expression}
	c.sub.Subscribe(topic, selector, c.handleMsg)
}

func (c *rocketQueue) handleMsg(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	defer rescue.Recover()
	for _, msg := range msgs {
		cctx, span := c.tracer.Start(ctx, "sub:"+msg.Topic, trace.WithSpanKind(trace.SpanKindConsumer))
		c.propagator.Inject(cctx, &MessageExtTextMapCarrier{msg: msg})
		span.SetAttributes(
			attribute.String("kafka.topic", msg.Topic),
			attribute.String("kafka.key", string(msg.GetKeys())),
		)
		if err := c.handler.Consume(ctx, msg.Topic, []byte(msg.GetKeys()), msg.Body); err != nil {
			span.RecordError(err)
			se := errors.FromError(err)
			log.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason)
			return consumer.ConsumeRetryLater, err
		}
	}
	return consumer.ConsumeSuccess, nil
}
