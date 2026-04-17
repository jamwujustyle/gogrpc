package main

import (
	"io"
	"log"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}
		log.Printf("receiving number: %d\n", req.Number)
		sum += req.Number
		count++
	}
}
