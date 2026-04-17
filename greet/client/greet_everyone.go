package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jamwujustyle/gogrpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "jam"},
		{FirstName: "aziza"},
		{FirstName: "shah"},
		{FirstName: "islam"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("error while receiving: %v\n", err)
			}

			log.Printf("received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
