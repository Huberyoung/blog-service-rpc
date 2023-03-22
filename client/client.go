package main

import (
	"blog_service_grpc/proto/pb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	ctx := context.Background()
	clientConn, err := GetClientConn(ctx, "localhost:8001", []grpc.DialOption{grpc.WithBlock()})
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "GO"})

	log.Printf("resp: %v err:%+v \n", resp, err)

}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
