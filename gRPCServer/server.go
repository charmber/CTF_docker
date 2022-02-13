package gRPCServer

import (
	"competition/gRPCServer/model"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func DockerGrpcServer() {
	grpcServer := grpc.NewServer()
	model.RegisterDockerServerServer(grpcServer, new(model.DockerServer))
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("监听失败", err)
		return
	} else {
		fmt.Println("success!")
	}
	_ = grpcServer.Serve(listen)
}
