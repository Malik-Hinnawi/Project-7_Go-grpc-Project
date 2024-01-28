package main

import pb "gRPC_basics/calculator/proto"

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	var k int32 = 2
	N := in.FirstNumber
	for N > 1 {
		if N%k == 0 {
			stream.SendMsg(&pb.PrimeResponse{Result: k})
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}
