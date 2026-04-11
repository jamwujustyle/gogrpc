package main

import (
	"context"
	"log"

	pb "github.com/jamwujustyle/gogrpc/app/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum func was invoked with: %v\n", in)
	sum:= in.NumOne + in.NumTwo

	return &pb.SumResponse{
			Result: sum,
		}, nil
}

