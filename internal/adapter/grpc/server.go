package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"my-grpc-go-server/internal/application"
	"my-grpc-go-server/internal/port"
	"net"
)

type GrpcAdapter struct {
	helloService application.HelloService
	grpcPort     int
	server       *grpc.Server
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
