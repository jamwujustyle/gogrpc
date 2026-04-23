package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/jamwujustyle/logger"
	pb "github.com/jamwujustyle/gogrpc/calculator/proto"
)

var addr string = "localhost:50053"

func main() {
	logger.InitLogger()

	conn, err:= grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error
		os.
	}
	defer conn.Close()

}