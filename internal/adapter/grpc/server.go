package grpc

import (
	"fmt"
	"github.com/lukmansjy/my-grpc-go-server/internal/port"
	"github.com/lukmansjy/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAdapter struct {
	helloService port.HelloServicePort
	grpcPort     int
	server       *grpc.Server
	hello.HelloServiceServer
}

func NewGrpcAdapter(helloService port.HelloServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		helloService: helloService,
		grpcPort:     grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %d : %v\n", a.grpcPort, err)
	}

	log.Printf("Server listen on port %d\n", a.grpcPort)

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	hello.RegisterHelloServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC on port %d : %d\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
