package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/alessiottavo/grpc-go/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	//insert credentials or else this will not work
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	r, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "David"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("%s\n", r.GetResult())
}
