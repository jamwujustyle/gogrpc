package main

import (
	pb "github.com/jamwujustyle/gogrpc/greet/proto" 
	"log"
	"context"
	"io"
	"fmt"
	)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("GreetManyTimes invoked")

	req:= &pb.GreetRequest{
		FirstName: "Zhamshid",
	}

	stream, err:= c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: \v\n", err)
	}

	for {
		msg, err:= stream.Recv()
		
		if err == io.EOF {
			break
		}
		
		fmt.Printf("GreetManyTimes: %s\n", msg.Result)
	}

}