package main

import (
	"context"
	"io"
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something wrong happened: %v\n", err)
		}

		log.Println(res)
	}
}
