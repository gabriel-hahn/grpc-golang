package main

import (
	"context"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	blog := &pb.Blog{
		AuthorId: "ana",
		Title:    "My Medium Post",
		Content:  "Content of the medium post",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
