package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt invoked with: %v\n", in)

	n := in.Number

	if n < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", n),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(n)),
	}, nil
}
