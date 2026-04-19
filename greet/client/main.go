package main

import (
	"log"
	"log/slog"

	pb "github.com/jamwujustyle/gogrpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true // can change that to false if needed
	opts := []grpc.DialOption{}

	if tls {
		cF := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(cF, "")

		if err != nil {
			slog.Info("error while loading CA trust cert", "err", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))

	}
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("failed to establish connection %v\n", err)
	}
	defer conn.Close()
	//...

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

	// doGreetManyTimes(c)

	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 2*time.Second)
}
