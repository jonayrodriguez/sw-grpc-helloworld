package helloworld

import (
	"context"
	"log"

	pb "github.com/logicnow/sw-grpc-helloworld/api/helloworld"
)

// Singleton pattern could be used here.

// Server struct for helloworld.
type Server struct {
	pb.UnimplementedHelloworldServer
}

// NewServer to create.
func NewServer() *Server {
	return &Server{}
}

// SayHelloworld implements helloworld.GreeterServer
func (s *Server) SayHelloworld(ctx context.Context, in *pb.HelloworldRequest) (*pb.HelloworldReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloworldReply{Message: "Hello World! " + in.GetName()}, nil
}
