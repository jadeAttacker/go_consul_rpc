package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "go_consul_rpc/helloworldProto"
	"go_consul_rpc/server/server_consul"

	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

var count int64

type server struct{} //服务对象

// SayHello 实现服务的接口 在proto中定义的所有服务都是接口
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) GetNum(ctx context.Context, in *pb.NumRequest) (*pb.NumReply, error) {
	return &pb.NumReply{Order: 1 + in.Id}, nil
}

func main() {
	go server_consul.RegisterServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()                  //起一个服务
	pb.RegisterGreeterServer(s, &server{}) // server 服务server对象
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
