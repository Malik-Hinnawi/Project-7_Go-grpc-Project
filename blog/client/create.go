package main

import (
	"context"
	pb "gRPC_basics/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked")
	blog := &pb.Blog{
		AuthorId: "Malik",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
