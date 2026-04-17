package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/jamwujustyle/gogrpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading: %v\n", err)
		}

		res := fmt.Sprintf("Hello %v", req.FirstName)

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error while sending data to client: %v\n", err)
		}

	}
}
