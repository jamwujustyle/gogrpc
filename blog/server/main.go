package main

import (
	"context"
	"log/slog"
	"net"
	"time"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
	"github.com/jamwujustyle/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50053"

type Server struct {
	pb.UnimplementedBlogServiceServer
}

func main() {
	logger.InitLogger()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		slog.Error("Failed to create mongo client", "err", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		slog.Error("Could not connect to mongo (Ping failed)", "err", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			slog.Error("Error disconnecting mongo", "err", err)
		}
	}()

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("Error establishing listener", "err", err)
	}
	slog.Info("Listening on", "addr", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		slog.Error("Failed to serve", "err", err)
	}

}
