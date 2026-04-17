package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	fmt.Println("doMax invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		fmt.Printf("error creating a stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Num: 1},
		{Num: 5},
		{Num: 3},
		{Num: 6},
		{Num: 2},
		{Num: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			fmt.Printf("sending a req: %v\n", req)
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
				fmt.Printf("error reading: %v\n", err)
			}

			fmt.Printf("current max: %v\n", res.Max)
		}
		close(waitc)
	}()

	<-waitc
}
