package main

import (
	"context"
	pb "gRPC_basics/calculator/proto"
	"log"
)

func doCalculate(c pb.CalculatorServiceClient) {
	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{FirstNumber: 1, SecondNumber: 2})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}
	log.Printf("Result: %v\n", res.Result)
}
