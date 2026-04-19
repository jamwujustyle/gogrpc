package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
)

var addr string = "localhost:50051"

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
	slog.Info("client starting")

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to establish connection: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doSum(c)
	// doPrimes(c)
	// doAverage(c)
	// doMax(c)
	doSqrt(c, -2)
}
