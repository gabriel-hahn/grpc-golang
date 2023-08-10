package main

import (
	"log"

	pb "github.com/gabriel-hahn/grpc-golang/blog/proto"
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

	c := pb.NewBlogServiceClient(conn)

	// createBlog(c)
	// readBlog(c, "id")
	// updateBlog(c, "id")
	// listBlog(c)
	deleteBlog(c, "id")
}
