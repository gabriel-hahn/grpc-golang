package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
)

func Greet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Gabriel"})
	if err != nil {
		log.Fatalf("Error on call Greet service: %v", err)
	}

	log.Print(res.Result)
}