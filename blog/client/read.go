package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
