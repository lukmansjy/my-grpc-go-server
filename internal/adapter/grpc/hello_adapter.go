package grpc

import (
	"context"
	"github.com/lukmansjy/my-grpc-proto/protogen/go/hello"
)

func (a *GrpcAdapter) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	greet := a.helloService.GenerateHello(req.Name)

	return &hello.HalloResponse{
		Greet: greet,
	}, nil
}
