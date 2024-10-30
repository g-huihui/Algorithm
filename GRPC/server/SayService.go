/**
 * @Author: Gong Yanhui
 * @Description: 服务端 protoc --go_out=../ --go-grpc_out=../ ./say.proto
 * @Date: 2024-10-30 14:25
 */

package main

import (
	"Algorithm/GRPC/api"
	"Algorithm/GRPC/common"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// SayServiceProvider 定义我们的服务
type SayServiceProvider struct {
	api.UnimplementedSayServer
}

// TestFunc 实现TestFunc方法
func (s *SayServiceProvider) TestFunc(ctx context.Context, req *api.TestRequest) (*api.TestResponse, error) {
	return &api.TestResponse{Message: "I say message: " + req.Name}, nil
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
	api.RegisterSayServer(grpcServer, &SayServiceProvider{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err: ", err)
	}
}
