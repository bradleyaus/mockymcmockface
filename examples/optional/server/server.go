// Don't edit, generated code

// Create the service types

//define method types

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/bradleyaus/mockymcmockface/pkg/responses"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"

	pb1 "github.com/bradleyaus/mockymcmockface/examples/optional"
)

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

	//For each service, the template will add one of these lines

	pb1.RegisterTestServer(s, &TestService{
		responseService: respService,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("failed when serving", err)
	}
}

type TestService struct {
	pb1.UnimplementedTestServer

	responseService *responses.ResponseService
}

func (t *TestService) Tester(ctx context.Context, req *pb1.Request) (*pb1.Response, error) {
	r, err := t.responseService.GetResponse("Test", "Tester", req)
	if err != nil {
		log.Panicln("Error when getting response", "Test", "Tester")
	}

	resp := &pb1.Response{}

	data, _ := json.Marshal(r.Data)
	err = jsonpb.Unmarshal(bytes.NewReader(data), resp)
	return resp, err
}
