package port

import (
	"context"

	"github.com/Siddhesh1992/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
)

// port is nothing all the service methods used or for client it would be the client service

type HelloClientPort interface {
	SayHello(ctx context.Context, in *hello.HelloRequest, opts ...grpc.CallOption) (*hello.HelloResponse, error)
	SayManyHellos(ctx context.Context, in *hello.HelloRequest, opts ...grpc.CallOption) (hello.HelloService_SayManyHellosClient, error)
}
