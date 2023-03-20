package server

import (
	"blog_service_grpc/pkg/bapi"
	"blog_service_grpc/proto/pb"
	"context"
	"encoding/json"
)

type TagServer struct {
	pb.UnimplementedTagServiceServer
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewApi("http://127.0.0.1:8080")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagListReply := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagListReply)
	if err != nil {
		return nil, err
	}
	return &tagListReply, nil
}

func (t *TagServer) mustEmbedUnimplementedTagServiceServer() {}
