package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"

	pb "github.com/jamwujustyle/gogrpc/blog/proto"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.UnimplementedBlogServiceServer
}

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String(a.Key, a.Value.Time().Format("15:04:05"))
			}
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				dir := filepath.Base(filepath.Dir(source.File))
				file := filepath.Base(source.File)
				return slog.String(a.Key, fmt.Sprintf("%s/%s:%d", dir, file, source.Line))
			}
			return a
		},
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		slog.Error("Error connecting to mongo", "err", err)
		os.Exit(1)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		slog.Error("Error establishing listener", "err", err)
		os.Exit(1)
	}
	slog.Info("Listening on port", "port", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	// reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		slog.Error("Failed to serve", "err", err)
		os.Exit(1)
	}

}
