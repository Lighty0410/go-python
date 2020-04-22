package main

import "os"

const (
	defaultJSONPath   = "data.json"
	defaultGRPCClient = "localhost:50051"
	defaultGRPCServer = ":8080"
	defaultHTTPServer = ":8081"

	jsonFilePath      = "JSON_FILE_PATH"
	grpcServerAddress = "GRPC_SERVER_ADDRESS"
	grpcClientAddress = "GRPC_CLIENT_ADDRESS"
	httpServerAddress = "HTTP_SERVER_ADDRESS"
)

type envs struct {
	jsonFile       string
	grpcServerAddr string
	grpcClientAddr string
	httpServerAddr string
}

func initEnv() envs {
	jsonFile := os.Getenv(jsonFilePath)
	if jsonFile == "" {
		jsonFile = defaultJSONPath
	}
	grpcServerAddr := os.Getenv(grpcServerAddress)
	if grpcServerAddr == "" {
		grpcServerAddr = defaultGRPCServer
	}
	grpcClientAddr := os.Getenv(grpcClientAddress)
	if grpcClientAddr == "" {
		grpcClientAddr = defaultGRPCClient
	}
	httpServerAddr := os.Getenv(httpServerAddress)
	if httpServerAddr == "" {
		httpServerAddr = defaultHTTPServer
	}
	return envs{
		jsonFile:       jsonFile,
		grpcServerAddr: grpcServerAddr,
		grpcClientAddr: grpcClientAddr,
		httpServerAddr: httpServerAddr,
	}
}
