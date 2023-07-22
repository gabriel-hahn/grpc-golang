package main

import (
	"context"

	pb "github.com/gabriel-hahn/grpc-golang/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	var result int32

	for _, value := range in.Number {
		result = result + value
	}

	return &pb.CalculatorResponse{
		Result: result,
	}, nil
}