package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	service "grpc_practice/hello-server/proto"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "test",
		"appkey": "test",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return true
}
func main() {
	creds, _ := credentials.NewClientTLSFromFile("D:\\Learn\\Practice\\grpc_practice\\key\\test.pem", "*.sends.com")
	//连接到server，禁用安全传输
	//coon, _ := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//安全传输
	//Tls+token双重验证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	coon, _ := grpc.Dial("127.0.0.1:9090", opts...)
	defer coon.Close()

	//建立连接
	client := service.NewSayHelloClient(coon)

	//rpc调用
	resp, _ := client.SayHello(context.Background(), &service.HelloRequest{RequestName: "sends"})
	fmt.Println(resp.GetResponseMsg())
}
