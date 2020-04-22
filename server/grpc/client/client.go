package client

import (
	api "test/server/grpc/api"

	"google.golang.org/grpc"
)

// NewGrpcClient creates a new grpc-client.
func NewGrpcClient(address string) (api.FilterTextClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return api.NewFilterTextClient(conn), nil
}
