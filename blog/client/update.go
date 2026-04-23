package main

import (
	"context"
	"log/slog"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	slog.Info("updateBlog invoked")

	req := &pb.Blog{
		Id:       id,
		AuthorId: "jam wind",
		Title:    "bitch",
		Content:  "hello world",
	}

	_, err := c.UpdateBlog(context.Background(), req)

	if err != nil {
		slog.Error("Error happened while updating", "err", err)
	}
	slog.Info("Updated!")

}
