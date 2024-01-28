package main

import (
	pb "gRPC_basics/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	var res int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading request in Max: %v\n", err)
		}
		if res < req.Number {
			res = req.Number
			err = stream.Send(&pb.MaxResponse{MaxNum: res})

			if err != nil {
				log.Fatalf("Error while sending max number: %v\n", err)
			}
		}

	}
	return nil
}
