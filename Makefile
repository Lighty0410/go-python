BIN_DIR := ./bin
PROTO_DIR := ./server/grpc/api/
PROTO_CLIENT := $(PROTO_DIR)/server.proto
PROTO_SERVER := $(PROTO_DIR)/client.proto

.PHONY:
build:
	go build -o $(BIN_DIR)

.PHONY:
test:
	go test -cover ./...

.PHONY:
clean:
	go clean
	rm -rf $(BIN_DIR)

.PHONY:
lint:
	golangci-lint run

.PHONY:
proto:
	protoc -I $(PROTO_DIR) $(PROTO_CLIENT) --go_out=plugins=grpc:$(PROTO_DIR)
	protoc -I $(PROTO_DIR) $(PROTO_SERVER) --go_out=plugins=grpc:$(PROTO_DIR)

mod:
	go mod tidy