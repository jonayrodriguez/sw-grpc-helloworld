package helloworld

import (
	"context"
	"log"
	"sync"

	pb "github.com/jonayrodriguez/sw-grpc-helloworld/api/helloworld"
)

// Singleton pattern could be used here.

// Server struct for helloworld.
type Server struct {
	pb.UnimplementedHelloworldServer
}

var instance *Server

// Call somente uma unica vez
var once sync.Once

// GetServerInstance to create or retrieve an existing instance
func GetServerInstance() *Server {
	// Using sync.Once guarantees the uniqueness of our instance, thread safe, ....
	once.Do(func() {
		instance = &Server{}
	})
	return instance
}

// SayHelloworld implementation
func (s *Server) SayHelloworld(ctx context.Context, in *pb.HelloworldRequest) (*pb.HelloworldReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloworldReply{Message: "Hello World! " + in.GetName()}, nil
}
