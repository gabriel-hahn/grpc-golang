package main

import (
	"fmt"
	"time"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d!", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
		time.Sleep(time.Second * 2)
	}

	return nil
}
