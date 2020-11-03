package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	hwpb "github.com/jonayrodriguez/sw-grpc-helloworld/api/helloworld"
	hwConfig "github.com/jonayrodriguez/sw-grpc-helloworld/internal/helloworld/config"
	hwServer "github.com/jonayrodriguez/sw-grpc-helloworld/internal/helloworld/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	defaultConfigFile = "./../config/default.yml"
)

var (
	version string = "v0.0.0"
)

func main() {

	/* Building binaries with version information and other metadata will improve your monitoring, logging, and debugging processes.
	   For example:
	   go build argsXYX -ldflags "-X main.Version=v1.0.0"
	*/
	fmt.Printf("Version: %s\n", version)

	var configFile string

	if len(os.Args) < 2 {
		configFile = defaultConfigFile
	} else {
		configFile = os.Args[1]

	}
	fmt.Printf("Configuration file used: %s\n", configFile)

	conf, err := hwConfig.LoadConfiguration(configFile)
	if err != nil {
		log.Fatalf("Configuration failure: %v", err)
	}

	address := fmt.Sprintf("%s%s%d", conf.Server.Host, ":", conf.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hwpb.RegisterHelloworldServer(s, hwServer.GetServerInstance())

	healthServer := health.NewServer()
	healthServer.SetServingStatus("helloword", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthServer)
	gracefulShutDown(s)
	fmt.Printf("Listening on %s\n", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

// Basic Channel to handle SIGINT and SIGTERM for a graceful shutdown
func gracefulShutDown(s *grpc.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer signal.Stop(ch)
		sig := <-ch
		errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
		log.Println(errorMessage)
		// Stop the service gracefully.
		s.GracefulStop()
	}()
}
