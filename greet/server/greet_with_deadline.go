package main

import (
	"context"
	"log/slog"
	"time"

	pb "github.com/jamwujustyle/gogrpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	slog.Info("GreetWithDeadline invoked with", "name", in)

	for range 3 {
		if ctx.Err() == context.DeadlineExceeded {
			errStr := "The client canceled the request"
			slog.Info(errStr)
			return nil, status.Error(codes.Canceled, errStr)
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
