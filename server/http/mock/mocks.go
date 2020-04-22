package mock

import (
	"context"
	"errors"
	api "test/server/grpc/api"
)

// Service represents a mock for the controller.
type Service struct{}

// GetTopics represents mock-func for the controller.
func (m *Service) GetTopics(word string) *api.Topics {
	return nil
}

// FilterText represents mock-func for the controller.
func (m *Service) FilterText(ctx context.Context, textCase string) (map[string]int64, error) {
	if textCase == "bad text" {
		return nil, errors.New("bad text")
	}
	return map[string]int64{
		"ok": 1,
	}, nil
}
