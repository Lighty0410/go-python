package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/controller"
	"test/server/grpc/client"
	grpcSv "test/server/grpc/server"
	"test/server/http"
	"test/utils"

	"github.com/valyala/fasthttp"
)

func main() {
	runServer()
}

func runServer() {
	envs := initEnv()

	jsonMap, err := utils.CreateJSONMap(envs.jsonFile)
	if err != nil {
		log.Fatalf("cannot create json map: %v", err)
	}
	grpcClient, err := client.NewGrpcClient(envs.grpcClientAddr)
	if err != nil {
		log.Fatalf("cannot create grpc client: %v", err)
	}
	ctrl := controller.New(jsonMap, grpcClient)
	grpcServer, err := grpcSv.NewGrpcServer(envs.grpcServerAddr, ctrl)
	if err != nil {
		log.Fatalf("cannot create grpc server: %v", err)
	}
	httpServer := http.NewHTTPServer(ctrl)
	fastServer := fasthttp.Server{Handler: httpServer.Run}
	go func() {
		err = fastServer.ListenAndServe(envs.httpServerAddr)
		if err != nil {
			log.Fatal(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	log.Printf("Shutting down the server due to the signal: %v\n", sig)
	grpcServer.GracefulStop()
	err = fastServer.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
