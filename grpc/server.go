package main

import (
	"context"
	"fmt"
	proto "go-learn/grpc/protoc"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedSayHelloServer
}

func (s *server) HelloWorld(ctx context.Context, req *proto.HelloRequest) (resp *proto.HelloReply, err error) {
	ctx.Done()
	return &proto.HelloReply{
		Message: "hello, " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:10084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterSayHelloServer(s, &server{})
	reflection.Register(s)
	fmt.Println("init service successfully")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
