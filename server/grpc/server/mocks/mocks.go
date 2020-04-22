package mocks

import (
	"context"
	grpc "test/server/grpc/api"
)

// MockTopic represents mock for the controller.
type MockTopic struct{}

// FilterText represents mock-func for the controller.
func (m *MockTopic) FilterText(ctx context.Context, textCase string) (map[string]int64, error) {
	return nil, nil
}

// GetTopics represents mock-func for the controller.
func (m *MockTopic) GetTopics(word string) *grpc.Topics {
	if word == "bad word" {
		return &grpc.Topics{
			Topics: []*grpc.Topic{
				{
					Key:   "bad word",
					Value: 69,
				},
			},
		}
	}
	return nil
}
