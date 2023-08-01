package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Sqrt(c pb.CalculatorServiceClient, n int32) {
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Code: %s. Error message from server: %s\n", e.Code(), e.Message())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
