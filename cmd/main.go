package main

import (
	mygrpc "github.com/lukmansjy/my-grpc-go-server/internal/adapter/grpc"
	app "github.com/lukmansjy/my-grpc-go-server/internal/application"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	hs := &app.HelloService{}

	grpcAdapter := mygrpc.NewGrpcAdapter(hs, 9090)

	grpcAdapter.Run()
}
