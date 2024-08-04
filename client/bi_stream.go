package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/sosweetham/grpc-demo/proto"
)

func callSayHelloBidiStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Print("bidi streaming started")
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break 
			}
			if err != nil {
				log.Fatalf("error while streaming: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request: %v", err)
		}
		time.Sleep(2*time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("bidi streaming finished")
}