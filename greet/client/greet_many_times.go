package main

import (
	"context"
	"io"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
)

func GreetManyTimes(c pb.GreetServiceClient) {
	streamRes, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{FirstName: "Gabriel"})

	if err != nil {
		log.Fatalf("Error on call Greet service: %v", err)
	}

	for {
		msg, err := streamRes.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
			break
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
