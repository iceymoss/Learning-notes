package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"StudyGin/jaeger_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type HelloSerivce struct{}

func (h *HelloSerivce) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {

	//接收context中的内容
	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		fmt.Println("get metadata err", ok)
	}
	for key, value := range md {
		fmt.Printf("%v: %v\n", key, value)
	}

	return &proto.HelloReply{
		Id:      "123456789",
		Request: req,
	}, nil
}
func (h *HelloSerivce) Ping(context.Context, *proto.IsEmpty) (*proto.Pong, error) {
	return nil, nil
}

func main() {
	//实例化server
	lit, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicln("监听失败", err)
	}
	//注册处理逻辑
	//NewServer创建一个未注册服务且尚未开始接受请求的 gRPC 服务器。
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &HelloSerivce{})
	log.Println(s.Serve(lit))

}
