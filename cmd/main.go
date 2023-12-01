package main

import (
	"log"
	"my-grpc-go-server/internal/adapter/grpc"
	"my-grpc-go-server/internal/application"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &application.HelloService{}

	grpcAdapter := grpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}
