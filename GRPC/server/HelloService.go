/**
 * @Author: Gong Yanhui
 * @Description: gRPC 服务端 protoc --go_out=../ --go-triple_out=../ ./hello.proto
 * @Date: 2024-10-26 16:57
 */

package main

import (
	"Algorithm/GRPC/api"
	"Algorithm/GRPC/common"
	"context"
	"github.com/dubbogo/grpc-go"
	"log"
	"net"
)

// HelloServiceProvider 定义我们的服务
type HelloServiceProvider struct {
	api.UnimplementedHelloServer
}

// SayHello 实现SayHello方法
func (s *HelloServiceProvider) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResp, error) {
	log.Println("[Server]: Receive " + req.Name)
	return &api.HelloResp{Message: "hello ,I'm received message: " + req.Name}, nil
}

func main() {

	// 监听本地端口
	listener, err := net.Listen(common.Network, common.Address)
	if err != nil {
		log.Panic("net.Listen err: ", err)
	}

	log.Println(common.Address + " net.Listing...")

	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 在gRPC服务器注册我们的服务
	api.RegisterHelloServer(grpcServer, &HelloServiceProvider{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err: ", err)
	}
}
