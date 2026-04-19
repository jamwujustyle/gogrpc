package main

import (
	"log"
	"log/slog"
	"net"
	"os"

	pb "github.com/jamwujustyle/gogrpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen  on: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true // change that to false if needed

	if tls {
		cF := "ssl/server.crt"
		kF := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(cF, kF)

		if err != nil {
			slog.Error("Failed loading certs", "err", err)
			os.Exit(1)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve")
	}
}
