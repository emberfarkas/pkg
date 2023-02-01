package rest

import (
	"context"

	"github.com/go-bamboo/pkg/middleware/logging"
	"github.com/go-bamboo/pkg/middleware/metadata"
	"github.com/go-bamboo/pkg/middleware/metrics"
	"github.com/go-bamboo/pkg/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Client = http.Client

func NewClient(endpoint string) (*Client, error) {
	middlewareChain := []middleware.Middleware{
		recovery.Recovery(),
		metadata.Client(),
		tracing.Client(),
		metrics.Client(),
		logging.Client(),
	}
	c, err := http.NewClient(context.TODO(), http.WithEndpoint(endpoint), http.WithMiddleware(middlewareChain...))
	if err != nil {
		return nil, err
	}
	return c, nil
}
