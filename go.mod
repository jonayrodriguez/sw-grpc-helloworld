module github.com/jonayrodriguez/sw-grpc-helloworld

go 1.15

require (
	github.com/cucumber/godog v0.10.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201030142918-24207fddd1c3 // indirect
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/validator.v2 v2.0.0-20200605151824-2b28d334fa05
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/jonayrodriguez/sw-grpc-helloworld/api/helloworld => ./api/helloworld
