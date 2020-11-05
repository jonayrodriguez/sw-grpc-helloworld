FROM golang:1.15

WORKDIR /go/sw-grpc-helloworld
COPY . .
RUN go get -d -v ./...
WORKDIR /go/sw-grpc-helloworld/cmd/helloworld

RUN go build main.go 

ENV SERVER_HOST=0.0.0.0
EXPOSE 50051

CMD ["./main"]

