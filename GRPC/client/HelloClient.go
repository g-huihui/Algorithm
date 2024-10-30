/**
 * @Author: Gong Yanhui
 * @Description: gRPC 客户端
 * @Date: 2024-10-29 15:38
 */

package main

import (
	"Algorithm/GRPC/api"
	"Algorithm/GRPC/common"
	"context"
	"github.com/dubbogo/grpc-go"
	"github.com/dubbogo/triple/pkg/triple"
	"log"
)

// 通过增加 --go-triple_out 参数生产的pb没有成功运行client端 triple.TripleConn 没有找到创建方式
func main() {

	// 连接服务器
	conn, err := grpc.Dial(common.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 如何将conn传入triple.TripleConn
	var tripleConn = triple.TripleConn{}

	// 建立gRPC连接
	grpcClient := api.NewHelloClient(&tripleConn)

	// 创建发送结构体
	req := api.HelloRequest{
		Name: "Taurus",
	}

	// 调用我们的服务(SayHello方法)
	// 同时传入了一个 context.Context, 在有需要时可以让我们改变RPC的行为, 比如超时/取消一个正在运行的RPC
	res, errAtt := grpcClient.SayHello(context.Background(), &req)
	if errAtt.GetError() != nil {
		log.Fatalf("grpcClient.SayHello err: %v", errAtt)
	}

	// 打印返回值
	log.Println(res)
}
