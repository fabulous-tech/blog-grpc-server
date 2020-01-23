package grpc

import (
	"context"
	pb "github.com/fabulous-tech/blog-proto/go/proto"
)

type Articles struct{}

func (a *Articles) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	rsp := new(pb.ListReply)
	return rsp, nil
}

func(a *Articles) Detail(ctx context.Context, req *pb.DetailRequest) (*pb.DetailReply, error) {
	rsp := new(pb.DetailReply)
	return rsp, nil
}