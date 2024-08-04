package main

import (
	"fmt"
	"log"

	pb "github.com/sosweetham/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// config, err := util.LoadConfig(".")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	var host string;
	var port string;

	fmt.Print("Enter host address: ")
	fmt.Scan(&host)
	fmt.Print("Enter port address: ")
	fmt.Scan(&port)

	hostAddress := host + ":" + port

	conn, err := grpc.NewClient(hostAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Soham", "Alice", "Bob"},
	}

	// callSayHello(client)

	callSayHelloServerStream(client, names)
}