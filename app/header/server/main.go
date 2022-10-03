package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	Name = "helloworld"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	info, _ := kratos.FromContext(ctx)
	if tr, ok := transport.FromServerContext(ctx); ok {
		tr.ReplyHeader().Set("app_name", info.Name())
	}

	return &helloworld.HelloReply{Message: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func main() {
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
	)
	httpSrv := http.NewServer(
		http.Address(":8000"),
	)

	s := &server{}

	helloworld.RegisterGreeterHTTPServer(httpSrv, s)
	helloworld.RegisterGreeterServer(grpcSrv, s)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
