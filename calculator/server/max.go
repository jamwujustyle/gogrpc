package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max invoked")

	var max int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("err: %v", err)
		}

		if num := req.Num; num > max {
			max = req.Num
			err = stream.Send(&pb.MaxResponse{
				Max: max,
			})
		}
		fmt.Printf("received: %d\n", req.Num)

		if err != nil {
			log.Printf("error while sending: %v\n", err)
		}
	}
}
