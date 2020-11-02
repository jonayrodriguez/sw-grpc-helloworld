// Package main implements a client for testing helloworld.
package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/jonayrodriguez/sw-grpc-helloworld/api/helloworld"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	address     = "localhost:50051"
	defaultName = "POC"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloworldClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t0 := time.Now()
	r, err := c.SayHelloworld(ctx, &pb.HelloworldRequest{Name: name})
	t1 := time.Now()
	log.Printf("The SayHelloworld request took %v to run.\n", t1.Sub(t0).Milliseconds())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	healthCheckConn := healthpb.NewHealthClient(conn)
	healthCheckResp, err := healthCheckConn.Check(ctx, &healthpb.HealthCheckRequest{})
	if err != nil {
		log.Fatalf("could not check the health: %v", err)
	}
	log.Printf("HealthCheck - check: %v", healthCheckResp)

}
