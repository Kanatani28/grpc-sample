package main

import (
	"log"
	"net"

	pb "../pb"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:55555")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	service := &incrementService{}

	pb.RegisterIncrementServiceServer(server, service)
	server.Serve(listen)
}
