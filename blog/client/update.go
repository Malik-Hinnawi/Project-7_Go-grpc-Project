package main

import (
	"context"
	pb "gRPC_basics/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--Update Blog was invocked--")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "New person",
		Title:    "A new title",
		Content:  "New content for the first blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
