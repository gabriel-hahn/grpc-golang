package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Gabriel"})
	if err != nil {
		log.Printf("Error on call Greet service: %v", err)
	}

	log.Print(res.Result)
}
