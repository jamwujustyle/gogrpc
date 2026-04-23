package main

import (
	"context"
	"log/slog"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	slog.Info("readBlog invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		slog.Error("Error happened while reading", "err", err)
	}

	slog.Info("Blog was read", "res", res)

	return res
}
