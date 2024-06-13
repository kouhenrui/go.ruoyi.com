package grpc

import (
	"fmt"
	"go.ruoyi.com/src/config"
	"go.ruoyi.com/src/internal/api/grpc/service/auth"
	pb "go.ruoyi.com/src/internal/api/proto/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 微服务文件生成：protoc --go_out=. --go-grpc_out=. auth.proto
func InitGrpc() {
	log.Println("开启监听tcp")
	// 加载证书和密钥
	//creeds, err := credentials.NewServerTLSFromFile("../../config/https/server.crt", "../../config/https/server.key")
	//if err != nil {
	//	log.Fatalf("Failed to generate credentials: %v", err)
	//
	//	panic(err)
	//}
	//grpc.Creds(creeds)
	// 创建 gRPC 服务器
	server := grpc.NewServer()

	// 将server结构体注册到grpc服务中
	pb.RegisterAuthServiceServer(server, &auth.AuthService{})
	ln, err := net.Listen("tcp", config.Tcp)
	if err != nil {
		fmt.Println("网络异常：", err)
		panic(err)
		//return
	}
	// 监听服务
	err = server.Serve(ln)
	if err != nil {
		fmt.Println("监听异常：", err)
		panic(err)
		//return
	}
	log.Println("tcp listen:", config.Tcp)
}
