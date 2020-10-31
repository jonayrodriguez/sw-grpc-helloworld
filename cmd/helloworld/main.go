package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	api "github.com/logicnow/sw-grpc-helloworld/api/helloworld"
	config "github.com/logicnow/sw-grpc-helloworld/internal/helloworld/config"
	server "github.com/logicnow/sw-grpc-helloworld/internal/helloworld/server"

	"google.golang.org/grpc"
)

const (
	defaultConfigPath = "../../config/default.yml"
)

var (
	version   string = "v0.0.0"
	buildTime string = time.Now().Format(time.RFC3339)
)

func main() {

	/* Building binaries with version information and other metadata will improve your monitoring, logging, and debugging processes.
	   For example:
	   go build argsXYX -ldflags "-X main.Version=v1.0.0 -X main.buildTime=$(date +"%Y.%m.%d.%H%M%S")"
	*/
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Build Time: %s\n", buildTime)

	var configPath string

	if len(os.Args) < 2 {
		configPath = defaultConfigPath
	} else {
		configPath = os.Args[1]

	}
	fmt.Printf("Configuration: %s\n", configPath)

	conf, err := config.LoadConfiguration(configPath)
	if err != nil {
		log.Fatalf("Configuration failure: %v", err)
	}

	address := fmt.Sprintf("%s%s%d", conf.Server.Host, ":", conf.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterHelloworldServer(s, server.NewServer())

	fmt.Printf("Listening on %s\n", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
