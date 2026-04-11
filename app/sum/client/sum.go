package main

import (
	"context"
	"log"

	pb "github.com/jamwujustyle/gogrpc/app/sum/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")
	res, err:= c.Sum(context.Background(), &pb.SumRequest{
		NumOne: 10,
		NumTwo: 15,
	})
	if err != nil {
		log.Fatalf("could not read: %v\n", err)
	}
	log.Printf("Sum: %v\n", res.Result)
}