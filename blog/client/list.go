package main

import (
	"context"
	"io"
	"log"
	"log/slog"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	slog.Info("listBlog invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		slog.Error("Error while calling ListBlogs", "err", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("Something happened", "err", err)
		}
		log.Println(res)
	}
}
