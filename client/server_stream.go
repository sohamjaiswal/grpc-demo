package main

import (
	"context"
	"io"
	"log"

	pb "github.com/sosweetham/grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("sending streamable data to server")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names.")
	}
	log.Printf("server streaming started...")
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("error occurred while streaming names: %v", err)
		}
		log.Println(message)
	}
	log.Printf("streaming finished")
}