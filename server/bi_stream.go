package main

import (
	"io"
	"log"

	pb "github.com/sosweetham/grpc-demo/proto"
	"google.golang.org/grpc/peer"
)

func (s *helloServer) SayHelloBidiStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {
	ctx := stream.Context()
	p, _ := peer.FromContext(ctx)
	log.Printf("recieved SayHelloClientStreaming invocation from: %s", p.Addr.String())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("%s's stream fininshed", p.Addr.String())
			return nil
		} else if err != nil {
			return err
		}
		log.Printf("got %s's request with name: %v",  p.Addr.String(), req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("failed to send res to %s: %v", p.Addr.String(), err)
			return err
		}
	}
}