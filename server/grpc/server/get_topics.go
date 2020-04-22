package server

import (
	"context"
	grpc "test/server/grpc/api"
)

// GrpcService is a grpc-server.
type GrpcService struct {
	Ctrl TopicService
}

// TopicService is a service for the controller.
type TopicService interface {
	GetTopics(word string) *grpc.Topics
}

// NewService creates a new grpc-server instance.
func NewService(ctrl TopicService) *GrpcService {
	return &GrpcService{
		Ctrl: ctrl,
	}
}

// GetTopics receive a word and returns a list of topics.
func (g *GrpcService) GetTopics(ctx context.Context, word *grpc.Word) (*grpc.Topics, error) {
	return g.Ctrl.GetTopics(word.Word), nil
}
