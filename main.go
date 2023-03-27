package main

import (
	"blog_service_grpc/proto/pb"
	"blog_service_grpc/server"
	"flag"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var grpcPort string
var httpPort string
var port string

func init() {
	//flag.StringVar(&grpcPort, "grpc_port", "8001", "grpc 启动端口号")
	//flag.StringVar(&httpPort, "http_port", "9001", "http 启动端口号")
	flag.StringVar(&port, "http_port", "8003", "启动端口号")
	flag.Parsed()
}

func RunTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})
	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	return s
}

func main() {
	l, err := RunTcpServer(port)
	if err != nil {
		log.Fatalf("Run Tcp Server err:%v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer()
	HttpS := RunHttpServer(port)

	go grpcS.Serve(grpcL)
	go HttpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		log.Fatalf("Run Server err:%v", err)
	}
}
