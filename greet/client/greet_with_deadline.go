package main

import (
	"context"
	"log/slog"
	"time"

	pb "github.com/jamwujustyle/gogrpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	slog.Info("doGreetWithDeadline invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "jam",
	}
	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			if e.Code() == codes.DeadlineExceeded {
				slog.Info("Deadline Exceeded")
			} else {
				slog.Error("gRPC error", "message", e.Message(), "code", e.Code())
			}
		}
		return
	}
	slog.Info("GreetWithDeadline", "res", res.GetResult())

}
