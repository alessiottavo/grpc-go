package main

import (
	"log"
	"net"

	pb "github.com/alessiottavo/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {

	var lis, err = net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to Listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
