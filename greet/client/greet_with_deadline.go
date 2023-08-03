package main

import (
	"context"
	"log"
	"time"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Gabriel",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			}

			log.Fatalf("Unexpected gRPC error: %v\n", err)
			return
		}

		log.Fatalf("A non gRPC error: %v\n", err)
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
