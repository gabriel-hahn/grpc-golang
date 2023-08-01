package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/calculator/proto"
)

func Sum(c pb.CalculatorServiceClient) {
	numbers := []int32{1, 44, 23}
	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{Number: numbers})
	if err != nil {
		log.Fatalf("Error on call Greet service: %v", err)
	}

	log.Print(res.Result)
}
