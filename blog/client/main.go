package main

import (
	"log/slog"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
	"github.com/jamwujustyle/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50053"

func main() {
	logger.InitLogger()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("error establishing connection", "err", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)

	readBlog(c, id)

	readBlog(c, "hello world")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)

}
