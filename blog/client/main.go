package main

import (
	"log/slog"
	"os"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
	"github.com/jamwujustyle/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50053"

func main() {
	logger.InitLogger()

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("error establishing connection", "err", err)
		os.Exit(1)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	createBlog(c)
}
