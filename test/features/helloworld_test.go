package helloworld

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/cucumber/godog"
	helloApi "github.com/jonayrodriguez/sw-grpc-helloworld/api/helloworld"
	helloServer "github.com/jonayrodriguez/sw-grpc-helloworld/internal/helloworld/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var listener *bufconn.Listener
var reply *helloApi.HelloworldReply

// Expected response prefix
const expectedResponsePrefix = "Hello World! "

// Given client is configured to contact server
func clientIsConfiguredToContactServer() error {

	listener = bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()
	helloApi.RegisterHelloworldServer(server, helloServer.GetServerInstance())

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

// When I say hello to server with "<message>"
func iSayHelloToServerWith(clientMessage string) error {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := helloApi.NewHelloworldClient(conn)
	reply, err = client.SayHelloworld(ctx, &helloApi.HelloworldRequest{Name: clientMessage})
	if err != nil {
		return fmt.Errorf("SayHelloworld failed: '%v'", err)
	}
	log.Printf("Response: %+v", reply)
	return nil
}

// Then server should respond with helloWorld "<message>"
func serverShouldRespondWithHelloWorld(clientMessage string) error {

	if reply.Message != expectedResponsePrefix+clientMessage {
		return fmt.Errorf("status code not as expected! Expected '%s', got '%s'", clientMessage, expectedResponsePrefix+clientMessage)
	}
	return nil
}

// Feature Steps
func FeatureContext(s *godog.Suite) {
	s.Step(`^client is configured to contact server$`, clientIsConfiguredToContactServer)
	s.Step(`^I say hello to server with "([^"]*)"$`, iSayHelloToServerWith)
	s.Step(`^server should respond with helloWorld "([^"]*)"$`, serverShouldRespondWithHelloWorld)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}
