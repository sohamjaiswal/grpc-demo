package main

import (
	"log"
	"strconv"

	"github.com/sosweetham/grpc-demo/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal(err)
	}

	hostAddress := config.Host + ":" + strconv.FormatInt(int64(config.Port), 10)

	// lis, err := net.Listen("tcp", hostAddress)

	// fmt.Printf("%+v\n", config)
}