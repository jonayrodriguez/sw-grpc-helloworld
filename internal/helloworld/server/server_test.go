package helloworld

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc"

	pb "github.com/logicnow/sw-grpc-helloworld/api/helloworld"
)

const (
	nameForTest               = "Test"
	expectedHelloworldMessage = "Hello World! " + nameForTest
)

type expectedResponse struct {
	actual   string
	expected string
}

var listener *bufconn.Listener

func init() {

	// Help you avoid starting up a service with a real port number, but still allowing testing of streaming RPCs.
	listener = bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterHelloworldServer(server, NewServer())

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func TestSayHelloworld(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloworldClient(conn)
	resp, err := client.SayHelloworld(ctx, &pb.HelloworldRequest{Name: "sdfds"})
	if err != nil {
		t.Fatalf("SayHelloworld failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	if resp.Message != expectedHelloworldMessage {
		t.Fatalf("Unexpected Response %+v", expectedResponse{resp.Message, expectedHelloworldMessage})
	}
}
