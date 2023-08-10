package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	blog := &pb.Blog{
		AuthorId: "ana",
		Title:    "My Medium Post Updated",
		Content:  "Content of the medium post updated",
	}

	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been updated")
}
