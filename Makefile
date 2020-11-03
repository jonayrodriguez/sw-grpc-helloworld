# TODO - Check and install dependencies if it's required.

# Note: You can use the -o flag to rename the executable or place it in a different location. 
# However, when building an executable for Windows and providing a different name, 
# be sure to explicitly specify the .exesuffix when setting the executable’s name

compile:
	@echo "Compiling proto files..."
	@protoc api\helloworld\helloworld.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
	@echo "Done"

build: clean compile test
	@echo "Bulding the application..."
	@go build -ldflags "-X main.version=v1.0.0" ./cmd/helloworld/main.go
	@echo "Done"

clean:
	@echo "Cleanning temp files..."
	@go clean -i ./...
	@echo "Done"

test: unitTest bddTest

unitTest: dependencies
	@echo "UnitTest running..."
	@go test -cpu 1,4 -timeout 7m .\cmd\...
	@go test -cpu 1,4 -timeout 7m .\internal\...

bddTest: dependencies
	@echo "BDD test running..."
	@cd .\test\features && godog .\helloworld.feature

dependencies:
	@go mod download

vendor:
	@go mod vendor	