package grpc

import (
	"net"

	pb "github.com/fabulous-tech/blog-proto/go/proto"
	"google.golang.org/grpc"
)

func Run() error{

	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterArticleServer(s, &Articles{})

	if err := s.Serve(l); err != nil {
		return err
	}
}