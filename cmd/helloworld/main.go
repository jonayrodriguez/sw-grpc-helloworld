package main

import (
	"log"
	"net"

	pb "github.com/logicnow/sw-grpc-helloworld/api/helloworld"
	hs "github.com/logicnow/sw-grpc-helloworld/internal/helloworld/server"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloworldServer(s, hs.NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
