package server

import (
	"context"
	grpc "test/server/grpc/api"
	"test/server/grpc/server/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTopics(t *testing.T) {
	tt := []struct {
		name          string
		word          *grpc.Word
		expectedTopic *grpc.Topics
	}{
		{
			name: "bad word",
			word: &grpc.Word{Word: "bad word"},
			expectedTopic: &grpc.Topics{
				Topics: []*grpc.Topic{
					{
						Key:   "bad word",
						Value: 69,
					},
				},
			},
		},
	}
	grpcService := NewService(&mocks.MockTopic{})
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			topic, _ := grpcService.GetTopics(context.Background(), tc.word)
			require.Equal(t, topic, tc.expectedTopic)
		})
	}
}

func TestNewService(t *testing.T) {
	ctrl := mocks.MockTopic{}
	grpcSerice := GrpcService{Ctrl: &ctrl}
	require.Equal(t, NewService(&ctrl), &grpcSerice)
}
