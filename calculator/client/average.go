package main

import (
	"context"
	"log"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Printf("doAverage invoked")

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("error while opening the stream: %v\n")
	}

	nums := []int32{3, 5, 9, 54, 23}

	for _, num := range nums {
		log.Printf("sending number: %d\n", num)
		stream.Send(&pb.AverageRequest{
			Number: num,
		})
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
