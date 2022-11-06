// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package status

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type StatusHTTPServer interface {
	Status(context.Context, *StatusRequest) (*StatusReply, error)
}

func RegisterStatusHTTPServer(s *http.Server, srv StatusHTTPServer) {
	r := s.Route("/")
	r.GET("/status", _Status_Status0_HTTP_Handler(srv))
}

func _Status_Status0_HTTP_Handler(srv StatusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/status.Status/Status")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Status(ctx, req.(*StatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StatusReply)
		return ctx.Result(200, reply)
	}
}

type StatusHTTPClient interface {
	Status(ctx context.Context, req *StatusRequest, opts ...http.CallOption) (rsp *StatusReply, err error)
}

type StatusHTTPClientImpl struct {
	cc *http.Client
}

func NewStatusHTTPClient(client *http.Client) StatusHTTPClient {
	return &StatusHTTPClientImpl{client}
}

func (c *StatusHTTPClientImpl) Status(ctx context.Context, in *StatusRequest, opts ...http.CallOption) (*StatusReply, error) {
	var out StatusReply
	pattern := "/status"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/status.Status/Status"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
