//go:unit

package main

import (
	"github.com/bradleyaus/mockymcmockface/pkg/server_template"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestGenerate(t *testing.T) {

	bytes, err := ioutil.ReadFile("../../examples/protoc-gen-data/input.dat")
	assert.NoError(t, err)

	var req pluginpb.CodeGeneratorRequest
	err = proto.Unmarshal(bytes, &req)
	assert.NoError(t, err)

	plugin := generate(&req, &server_template.Options{
		TemplateName: "server.tmpl",
		TemplatePath: "../../server.tmpl",
	})

	_, err = proto.Marshal(plugin.Response())
	assert.NoError(t, err)
	//TODO: Test snapshots could be useful
}


//protoc --plugin=mockymock=build/protoc-gen-mockymock --mockymock_out=examples -I examples/basic/protobuf examples/basic/protobuf/basic.proto --go-grpc_out=examples/basic/grpc --go_out=examples/basic/grpc