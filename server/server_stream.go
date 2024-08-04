package main

import (
	"log"
	"time"

	pb "github.com/sosweetham/grpc-demo/proto"
	"google.golang.org/grpc/peer"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList,
	stream pb.GreetService_SayHelloServerStreamingServer) error {
	ctx := stream.Context()
	p, _ := peer.FromContext(ctx)
	log.Printf("recieved SayHelloServerStreaming invocation from: %s", p.Addr.String())
	log.Printf("got req w/ names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2*time.Second)
	}
	return nil
}