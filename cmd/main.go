package main

import (
	"context"
	"log"

	"github.com/Siddhesh1992/my-grpc-go-client/internal/adapter/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(log.Writer())

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:9090", opts...)

	if err != nil {
		log.Fatalln("can not connect to grpc server :", err)
	}

	defer conn.Close()

	helloAdapter, err := hello.NewHelloAdapter(conn)

	if err != nil {
		log.Fatalln("Can not create HelloAdapter :", err)
	}

	runSayHello(helloAdapter, "Bruce Wayne")
}

func runSayHello(adapter *hello.HelloAapter, name string) {
	great, err := adapter.SayHello(context.Background(), name)

	if err != nil {
		log.Fatalln("err not say hello :", err)
	}

	log.Println(great.Greet)
}
