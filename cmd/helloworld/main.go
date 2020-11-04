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

var (
	version    string = "v0.0.0"
	configFile string
)

func main() {

	/* Building binaries with version information and other metadata will improve your monitoring, logging, and debugging processes.
	   For example:
	   go build argsXYX -ldflags "-X main.Version=v1.0.0"
	*/
	log.Printf("Version: %s\n", version)

	if len(os.Args) >= 2 {
		configFile = os.Args[1]

	}
	log.Println("Loading configuration ...")
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

	// HealhCheck
	healthServer := health.NewServer()
	healthServer.SetServingStatus("helloword", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthServer)

	// Gracefull Shutdown
	gracefulShutDown(s)

	log.Printf("Listening on %s\n", address)
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
		s.GracefulStop()
	}()
}
