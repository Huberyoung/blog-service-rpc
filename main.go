package main

import (
	"blog_service_grpc/proto/pb"
	"blog_service_grpc/server"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8001", "启动端口号")
	flag.Parsed()
}

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("net.Listen err: %v", err)
	}

	err = s.Serve(listen)
	if err != nil {
		log.Printf("s.Serve err: %v", err)
	}
}
