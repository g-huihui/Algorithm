/**
 * @Author: Gong Yanhui
 * @Description: 客户端
 * @Date: 2024-10-30 14:53
 */

package main

import (
	"Algorithm/GRPC/api"
	"Algorithm/GRPC/common"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	// 连接服务器
	conn, err := grpc.NewClient(common.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}

	// 建立gRPC连接
	grpcClient := api.NewSayClient(conn)

	// 创建发送结构体
	req := api.TestRequest{
		Name: "Taurus",
	}

	// 调用我们的服务(TestFunc方法)
	// 同时传入了一个 context.Context, 在有需要时可以让我们改变RPC的行为, 比如超时/取消一个正在运行的RPC
	res, err := grpcClient.TestFunc(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpcClient.TestFunc err: %v", err)
	}

	// 打印返回值
	log.Println(res)
}
