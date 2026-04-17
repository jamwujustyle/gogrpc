package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while readming stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}
}
