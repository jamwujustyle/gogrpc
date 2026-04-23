package main

import (
	"context"
	"log/slog"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	slog.Info("deleteBlog invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		slog.Error("Error while deleting", "err", err)
	}

	slog.Info("Blog was deleted")

}
