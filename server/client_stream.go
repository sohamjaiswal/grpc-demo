package main

import (
	"io"
	"log"

	pb "github.com/sosweetham/grpc-demo/proto"
	"google.golang.org/grpc/peer"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	ctx := stream.Context()
	p, _ := peer.FromContext(ctx)
	log.Printf("recieved SayHelloClientStreaming invocation from: %s", p.Addr.String())
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("%s's client stream fininshed, sending response: %v", p.Addr.String(), messages)
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		} else if err != nil {
			return err
		}
		log.Printf("got %s's request with name: %v",  p.Addr.String(), req.Name)
		messages = append(messages, "Hello " + req.Name)
	}
}