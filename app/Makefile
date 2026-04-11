BIN_DIR = bin
APP_DIR = sum
PROTO_DIR = $(APP_DIR)/proto
SERVER_DIR = $(APP_DIR)/server
CLIENT_DIR = $(APP_DIR)/client

PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

.DEFAULT_GOAL := build


build:
	protoc -I=. \
		--go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto
	go build -o $(BIN_DIR)/server .$(SERVER_DIR)
	go build -o $(BIN_DIR)/client .$(CLIENT_DIR)

server:
	.${BIN_DIR}/server
client:
	.${BIN_DIR}/client

cleanup:
	rm -rf $(PROTO_DIR)/*.pb.go
