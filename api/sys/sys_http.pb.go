// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package sys

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

type SysHTTPServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResp, error)
	CheckResource(context.Context, *CheckResourceRequest) (*CheckResourceReply, error)
}

func RegisterSysHTTPServer(s *http.Server, srv SysHTTPServer) {
	r := s.Route("/")
	r.GET("/x/internal/sys/auth", _Sys_Auth0_HTTP_Handler(srv))
	r.GET("/x/internal/sys/checkcasbin", _Sys_CheckResource0_HTTP_Handler(srv))
}

func _Sys_Auth0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.sys.Sys/Auth")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Auth(ctx, req.(*AuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthResp)
		return ctx.Result(200, reply)
	}
}

func _Sys_CheckResource0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.sys.Sys/CheckResource")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckResource(ctx, req.(*CheckResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckResourceReply)
		return ctx.Result(200, reply)
	}
}

type SysHTTPClient interface {
	Auth(ctx context.Context, req *AuthRequest, opts ...http.CallOption) (rsp *AuthResp, err error)
	CheckResource(ctx context.Context, req *CheckResourceRequest, opts ...http.CallOption) (rsp *CheckResourceReply, err error)
}

type SysHTTPClientImpl struct {
	cc *http.Client
}

func NewSysHTTPClient(client *http.Client) SysHTTPClient {
	return &SysHTTPClientImpl{client}
}

func (c *SysHTTPClientImpl) Auth(ctx context.Context, in *AuthRequest, opts ...http.CallOption) (*AuthResp, error) {
	var out AuthResp
	pattern := "/x/internal/sys/auth"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.sys.Sys/Auth"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysHTTPClientImpl) CheckResource(ctx context.Context, in *CheckResourceRequest, opts ...http.CallOption) (*CheckResourceReply, error) {
	var out CheckResourceReply
	pattern := "/x/internal/sys/checkcasbin"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.sys.Sys/CheckResource"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
