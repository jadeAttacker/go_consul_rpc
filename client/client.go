package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"go_consul_rpc/client/client_consul"
	pb "go_consul_rpc/helloworldProto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	defaultId   = int32(32)
)

var lastIndex uint64

func main() {
	addr2 := client_consul.Lookup()
	// 建立链接
	conn, err := grpc.Dial(addr2.(string), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	id2 := defaultId
	if len(os.Args) > 2 {
		name = os.Args[1]
		id, _ := strconv.Atoi(os.Args[2])
		id2 = int32(id)
	}
	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r1, err1 := c.GetNum(ctx, &pb.NumRequest{Id: id2})
	if err != nil {
		log.Fatalf("could not greet: %v", err1)
	}
	log.Printf("Order is: %v", r1.Order)
}
