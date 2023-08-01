package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
)

func GreetEveryone(c pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{FirstName: "Gabriel"},
		{FirstName: "Ana"},
		{FirstName: "Paulo"},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error on call Greet service: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)

			stream.Send(req)
			time.Sleep(2 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading the stream: %v", err)
				break
			}

			log.Printf("GreetEveryone: %s\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
