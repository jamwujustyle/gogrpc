package main

import (
	"fmt"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	fmt.Printf("Primes invoked with: %v\n", in)

	n := in.Number
	d := int64(2)

	for n > 1 {
		if n%d == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: d,
			})
			n /= d
		} else {
			d++
		}
	}
	return nil
}
