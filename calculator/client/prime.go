package main

import (
	"context"
	pb "gRPC_basics/calculator/proto"
	"io"
	"log"
)

func doPrime(c pb.CalculatorServiceClient) {
	req := &pb.PrimeRequest{FirstNumber: 120}

	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling doPrime: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Result: %v\n", msg.Result)
	}
}
