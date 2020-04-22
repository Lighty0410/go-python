package server

import (
	"log"
	"net"
	api "test/server/grpc/api"

	"google.golang.org/grpc"
)

// NewGrpcServer creates a new grpc-server.
func NewGrpcServer(address string, topicService TopicService) (*grpc.Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	service := NewService(topicService)
	server := grpc.NewServer()
	api.RegisterGetTopicsServer(server, service)
	go func() {
		err = server.Serve(listener)
		if err != nil {
			log.Println(err)
		}
	}()
	return server, nil
}
