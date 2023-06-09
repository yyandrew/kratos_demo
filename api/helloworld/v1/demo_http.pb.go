// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v4.22.3
// source: helloworld/v1/demo.proto

package v1

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

const OperationDemoListDemo = "/api.helloworld.Demo/ListDemo"

type DemoHTTPServer interface {
	ListDemo(context.Context, *ListDemoRequest) (*ListDemoReply, error)
}

func RegisterDemoHTTPServer(s *http.Server, srv DemoHTTPServer) {
	r := s.Route("/")
	r.GET("/demos/list", _Demo_ListDemo0_HTTP_Handler(srv))
}

func _Demo_ListDemo0_HTTP_Handler(srv DemoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDemoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoListDemo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDemo(ctx, req.(*ListDemoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDemoReply)
		return ctx.Result(200, reply)
	}
}

type DemoHTTPClient interface {
	ListDemo(ctx context.Context, req *ListDemoRequest, opts ...http.CallOption) (rsp *ListDemoReply, err error)
}

type DemoHTTPClientImpl struct {
	cc *http.Client
}

func NewDemoHTTPClient(client *http.Client) DemoHTTPClient {
	return &DemoHTTPClientImpl{client}
}

func (c *DemoHTTPClientImpl) ListDemo(ctx context.Context, in *ListDemoRequest, opts ...http.CallOption) (*ListDemoReply, error) {
	var out ListDemoReply
	pattern := "/demos/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoListDemo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
