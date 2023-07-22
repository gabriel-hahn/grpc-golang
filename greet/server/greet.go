package main

import (
	"context"
	"fmt"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	greetStr := fmt.Sprintf("Hello %s", in.FirstName)

	return &pb.GreetResponse{
		Result: greetStr,
	}, nil
}
