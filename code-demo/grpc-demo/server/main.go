package main

import (
	"code-demo/grpc-demo/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

//实现proto接口

type HelloServer struct{}

func (h *HelloServer) SayHello(ctx context.Context, req *proto.HolleRequest) (*proto.HelloReply, error) {
	t := strconv.Itoa(int(time.Now().Unix()))
	res := &proto.HelloReply{
		Id:      "grpc" + t,
		Request: req,
	}
	return res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("监听服务失败")
		return
	}

	s := grpc.NewServer()

	proto.RegisterGreeterServer(s, &HelloServer{})

	err = s.Serve(listen)
	if err != nil {
		log.Println("服务启动失败", err)
	}
}
