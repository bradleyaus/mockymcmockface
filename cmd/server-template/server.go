package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/bradleyaus/mockymcmockface/pkg/responses"
	pb "github.com/bradleyaus/mockymcmockface/examples/basic"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"log"
	"net"
)



// This is just a file which is kept up to date with the server_template
// 1. Makes it easier to understand what the server_template will do
// 2. Makes it easier to examples ideas related to the server_template

// For each parser, the server_template will add one of these types
type TestService struct {
	pb.UnimplementedTestServer

	responseService *responses.ResponseService
}

func main() {

	respService := responses.NewService()

	//TODO: Param the address
	addr := "localhost:5051"
	log.Println("Listening on", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("failed when listening", err)
	}

	s := grpc.NewServer()

	//For each parser, the server_template will add one of these lines
	pb.RegisterTestServer(s, &TestService{
		responseService:         respService,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed when serving", err)
	}
}

// For each method in each parser, the server_template will add one of these methods
func (t *TestService) Tester(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	r, err := t.responseService.GetResponse("Test", "Tester", req)
	if err != nil {
		log.Panicln("Error when getting response", "Test", "Tester")
	}

	resp = &pb.Response{}

	data, _ := json.Marshal(r.Data)
	err = jsonpb.Unmarshal(bytes.NewReader(data), resp)
	return resp, err
}

func (t *TestService) Tester2(context.Context, *pb.Request1) (resp *pb.Response1, err error) {
	resp = &pb.Response1{}

	err = jsonpb.Unmarshal(bytes.NewReader([]byte("examples")), resp)
	return resp, nil
}
