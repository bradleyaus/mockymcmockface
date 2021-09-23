//go:unit

package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestGenerate(t *testing.T) {

	bytes, err := ioutil.ReadFile("../../test/protoc-gen-data/input.dat")

	assert.NoError(t, err)

	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(bytes, &req)

	generate(&req)
}
