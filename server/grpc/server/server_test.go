package server

import (
	"net"
	"testing"
	"time"

	"test/server/grpc/server/mocks"

	"github.com/stretchr/testify/require"
)

func TestNewGrpcServer(t *testing.T) {
	server, err := NewGrpcServer(":8081", &mocks.MockTopic{})
	require.NoError(t, err)
	server.Stop()
	time.Sleep(time.Millisecond * 50)
	_, err = net.Listen("tcp", ":8081")
	require.NoError(t, err)
	_, err = NewGrpcServer(":8081", &mocks.MockTopic{})
	require.EqualError(t, err, "listen tcp :8081: bind: address already in use")
}
