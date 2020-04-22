package mocks

import (
	"context"
	"errors"
	api "test/server/grpc/api"

	"google.golang.org/grpc"
)

// MockClient represents mock for the controller.
type MockClient struct{}

// FilterText represents mock func for the controller.
func (m *MockClient) FilterText(ctx context.Context, text *api.Text, opts ...grpc.CallOption) (*api.Answer, error) {
	if text.Text == "bad text" {
		return nil, errors.New("unacceptable text")
	}
	okAnswer := make(map[string]int64)
	okAnswer["ok"] = 1
	return &api.Answer{Result: okAnswer}, nil
}
