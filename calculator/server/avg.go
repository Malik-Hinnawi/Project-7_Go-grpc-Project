package main

import (
	pb "gRPC_basics/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				AvgResult: sum / count,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		count++
		sum += float64(req.Number)
	}
	return nil
}
