package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to establish connection: %v\n", err)
	}
	fmt.Printf("listening on: %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
}
