package main

import (
	"code-demo/grpc-demo/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println("连接失败")
	}

	defer conn.Close()

	client := proto.NewGreeterClient(conn)

	res, err := client.SayHello(context.Background(), &proto.HolleRequest{
		Name:    "SayHello",
		Url:     "/rpc",
		Gender:  1,
		M:       map[string]string{"name": "iceymoss", "age": "18"},
		AddTime: nil,
	})
	if err != nil {
		log.Println("req fail: SayHello")
	}

	fmt.Println("请求结果:", res)

}
