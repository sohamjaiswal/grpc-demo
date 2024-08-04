package main

import (
	"log"
	"net"
	"strconv"

	pb "github.com/sosweetham/grpc-demo/proto"
	"github.com/sosweetham/grpc-demo/util"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal(err)
	}

	hostAddress := config.Host + ":" + strconv.FormatInt(int64(config.Port), 10)

	lis, err := net.Listen("tcp", hostAddress)

	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Greet Service registered at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	// fmt.Printf("%+v\n", config)
}