package main

import (
	"context"
	pb "gRPC_basics/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("doGreetWithDeadline was incoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Malik",
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded !")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}

		} else {
			log.Fatalf("A non gRPC error %v\n", err)
		}
	}
	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
