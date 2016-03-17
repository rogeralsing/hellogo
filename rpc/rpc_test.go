package rpc

import (
	"testing"
	"net"
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/stretchr/testify/assert"
	"github.com/golang/protobuf/proto"
)

func TestServerConnection(t *testing.T) {
	const address = "127.0.0.1:8090"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	server := HelloServiceServerConcrete{}
	RegisterHelloServiceServer(s,&server)
	go s.Serve(lis)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	client := NewHelloServiceClient(conn)
	response,err := client.SayHello(context.Background(), &HelloRequest{Greeting: proto.String("hello")})
	assert.Equal(t,"hej",*response.Reply)
}