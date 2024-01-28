package main

import (
	"context"
	pb "gRPC_basics/calculator/proto"
	"log"
)

func doAvg(c pb.CalculatorServiceClient) {
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg %v\n", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving the result from Avg: %v\n", err)
	}

	log.Printf("The average is: %v\n", res)
}
