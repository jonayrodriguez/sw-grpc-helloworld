# HELLOWORLD (Go/GRPC) POC


This POC is designed to get you up and running with a project structure optimized for developing
gRPC services in Go. It promotes the best practices that follow the [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
and [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). 



## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.13 or above**.

After installing Go, run the following commands to start experiencing this starter kit:

```shell
# download the POC
git clone https://github.com/jonayrodriguez/sw-grpc-helloworld.git

# Pulling all dependencies
cd sw-grpc-helloworld
go mod vendor

# start the server (This will hold your console waiting for requests).
go run cmd/helloworld/main.go

# start the client in a new console for testing for now (this is just to show how it will be used).
go run cmd/helloworld_client/main.go

```

## Project Layout

The POC uses the following project layout:
 
```
.
├── api                    OpenAPI/Swagger specs, JSON schema files, protocol definition files
│   └── helloworld   	   Helloworld api definition
├── cmd                    Main applications of the project
│   └── helloworld   	   Helloworld main go file
│   └── helloworld_client  Helloworld client for testing
├── config                 Configuration files yaml,...
├── internal               Private application and library code
    ├── helloworld         Private folder for helloworld application
        ├── server         Helloworld server
        ├── config         Helloworld configuration library
├── test               	   Artifacts used to test application
    ├── features           BDD test scenarios and scripts
    
```

**NOTE: Keep in mind that this documentation will be updated when more features are added (folder : pkg, test, docs, etc.)**


The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html).

Within each feature, code are organized in layers, following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## BDD

The project uses [godog](https://github.com/cucumber/godog) to implement BDD test scenarios.
See [installation instructions](https://github.com/cucumber/godog#install) to setup your environment.

Run the following commands to execute the BDD tests.

```shell

# Navigate to the "features" folder:
cd ./test/features

# Run the hellowworld feature tests:
godog helloworld.feature

```
