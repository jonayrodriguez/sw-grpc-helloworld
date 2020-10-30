module github.com/logicnow/sw-grpc-helloworld

go 1.15

require (
	github.com/golang/protobuf v1.4.3 // indirect
	golang.org/x/net v0.0.0-20201029221708-28c70e62bb1d // indirect
	golang.org/x/sys v0.0.0-20201029080932-201ba4db2418 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201030160639-589136c8afd9 // indirect
	google.golang.org/genproto v0.0.0-20201029200359-8ce4113da6f7 // indirect
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

replace github.com/logicnow/sw-grpc-helloworld/api/helloworld => ./api/helloworld
