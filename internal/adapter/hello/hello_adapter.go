package hello

import (
	"context"
	"log"

	"github.com/Siddhesh1992/my-grpc-go-client/internal/port"
	"github.com/Siddhesh1992/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
)

type HelloAapter struct {
	helloClient port.HelloClientPort
}

func NewHelloAdapter(conn *grpc.ClientConn) (*HelloAapter, error) {
	client := hello.NewHelloServiceClient(conn)
	return &HelloAapter{
		helloClient: client,
	}, nil
}

func (a *HelloAapter) SayHello(ctx context.Context, name string) (*hello.HelloResponse, error) {
	helloRequest := &hello.HelloRequest{
		Name: name,
	}

	greet, err := a.helloClient.SayHello(ctx, helloRequest)

	if err != nil {
		log.Fatalln("Error on SayHello: %v", err)
	}

	return greet, nil
}
