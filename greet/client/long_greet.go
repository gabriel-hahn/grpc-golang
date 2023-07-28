package main

import (
	"context"
	"log"
	"time"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
)

func LongGreet(c pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{FirstName: "Gabriel"},
		{FirstName: "Ana"},
		{FirstName: "Paulo"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error on call Greet service: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error on receiving response from stream: %v", err)
	}

	log.Printf("LongGreet response: %s\n", res.Result)
}
