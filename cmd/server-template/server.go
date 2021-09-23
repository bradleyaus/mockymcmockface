package server_template

import (
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/bradleyaus/mockymcmockface/test/basic/grpc"
)

//This is just a file which is kept up to date with the template to make it easier to update the template with new features

type TestService struct {

}

func main() {
	//TODO: Param the address
	lis, err := net.Listen("TCP", "localhost:5051")
	if err != nil {
		log.Fatal("failed when listening", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestServer(s, &TestService{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed when serving", err)
	}
}

func (t *TestService) Tester(context.Context, *pb.Request) (*pb.Response, error) {


	msg := proto.Message{}

	res := protoiface.MessageV1()
	return &interface{}, nil
}

Tester2(context.Context, *pb.Request1) (*Response1, error)
