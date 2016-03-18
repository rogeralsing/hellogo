package rpc

import (
	"golang.org/x/net/context"
	"github.com/golang/protobuf/proto"
)

type Server struct {
}

func (this Server) SayHello(context context.Context,request *HelloRequest) (*HelloResponse, error) {
	response := &HelloResponse{
		Reply: proto.String("hej"),
	}
	return response,nil
}