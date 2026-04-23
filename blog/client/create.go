package main

import (
	"context"
	"log/slog"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	slog.Info("---createBlog invoked---")

	blog := &pb.Blog{
		AuthorId: "jam wind",
		Title:    "hello world",
		Content:  "some content here",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		slog.Error("unexpected error", "err", err)
	}

	slog.Info("response from server", "res", res.GetId())
	return res.GetId()
}
