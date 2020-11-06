package test

import (
	"context"
	"flag"
	"fmt"
	"github.com/cucumber/godog/colors"
	"log"
	"net"
	"os"
	"testing"

	"github.com/cucumber/godog"
	helloApi "github.com/jonayrodriguez/sw-grpc-helloworld/api/helloworld"
	helloServer "github.com/jonayrodriguez/sw-grpc-helloworld/internal/helloworld/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

var listener *bufconn.Listener
var reply *helloApi.HelloworldReply

// Expected response prefix
const expectedResponsePrefix = "Hello World! "

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name: "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

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

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		log.Println("BeforeSuite") })
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		log.Println("Before Scenario") // clean the state before every scenario
	})

	ctx.Step(`^client is configured to contact server$`, clientIsConfiguredToContactServer)
	ctx.Step(`^I say hello to server with "([^"]*)"$`, iSayHelloToServerWith)
	ctx.Step(`^server should respond with helloWorld "([^"]*)"$`, serverShouldRespondWithHelloWorld)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}
