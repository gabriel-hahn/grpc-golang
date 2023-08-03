package main

import (
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const addr string = "localhost:50051"

func main() {
	opts := []grpc.DialOption{}
	certFile := "ssl/ca.crt"

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatalf("Error while loading CA trust certificate: %v\n", err)
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	Greet(c)
	// GreetManyTimes(c)
	// LongGreet(c)
	// GreetEveryone(c)
	// GreetWithDeadline(c, 1*time.Second)
}
