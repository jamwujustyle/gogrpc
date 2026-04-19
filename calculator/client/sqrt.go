package main

import (
	"context"
	"log/slog"
	"os"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	slog.Info("doSqrt invoked", "number", n)
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		if e, ok := status.FromError(err); ok {
			slog.Warn("gRPC error",
				"message", e.Message(),
				"code", e.Code().String(),
			)
			if e.Code() == codes.InvalidArgument {
				slog.Debug("invalid input")
			}
		} else {
			slog.Error("infrastructure or network failure", "err", err)
			os.Exit(1)
		}
		return
	}
	slog.Info("sqrt result received", "result", res.GetResult())
}
