package main

import (
	"context"
	"log"

	pb "github.com/sosweetham/grpc-demo/proto"
	"google.golang.org/grpc/peer"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	p, _ := peer.FromContext(ctx)
	log.Printf("recieved SayHello invocation from: %s", p.Addr.String())
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}