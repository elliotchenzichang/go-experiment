package main

import (
	"context"
	"fmt"
	"log"
	"time"

	proto "go-learn/grpc/protoc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewSayHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.HelloWorld(ctx, &proto.HelloRequest{Name: "World"})

	if err != nil {
		log.Fatalf("could not get the correct result, err: %s", err)
		return
	}

	fmt.Println(result.Message)

}
