package main

import (
	"log"
	"time"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// Greet(c)
	// GreetManyTimes(c)
	// LongGreet(c)
	// GreetEveryone(c)
	GreetWithDeadline(c, 1*time.Second)
}
