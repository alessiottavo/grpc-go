package main

import (
	"context"
	"log"

	pb "github.com/alessiottavo/grpc-go/greet/proto"
)

// Unary API client implementation
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
