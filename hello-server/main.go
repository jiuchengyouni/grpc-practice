package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	service "grpc_practice/hello-server/proto"
	"net"
)

type server struct {
	service.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("无token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "test" || appKey != "test" {
		return nil, errors.New("token错误")
	}
	return &service.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//TLS认证
	creds, _ := credentials.NewServerTLSFromFile("D:\\Learn\\Practice\\grpc_practice\\key\\test.pem", "D:\\Learn\\Practice\\grpc_practice\\key\\test.key")
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//在grpc服务端中去注册我们自己编写的服务
	service.RegisterSayHelloServer(grpcServer, &server{})
	err := grpcServer.Serve(listen)
	if err != nil {
		return
	}
}
