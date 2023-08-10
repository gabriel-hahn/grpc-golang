package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	blog := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was deleted")
}
