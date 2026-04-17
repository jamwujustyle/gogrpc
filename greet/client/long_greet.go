package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jamwujustyle/gogrpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Jam"},
		{FirstName: "Islam"},
		{FirstName: "Aziza"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", reqs)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from LongGreet %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
