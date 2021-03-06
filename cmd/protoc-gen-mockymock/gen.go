package main

import (
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	// This is a compiler that is called by protoc, which means we have to compile this into a binary and run.
	// It passes all that is required for compilation through stdin
	// The problem with that is it's difficult to quickly iterate on a solution
	// So to make development easier/allow debugging, we can save the input from stdin to a file and then use that in a test
	os.MkdirAll("test/protoc-gen-data/", os.ModePerm)
	ioutil.WriteFile("test/protoc-gen-data/input.dat", input, os.ModePerm)

	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	generate(&req)
}

func generate(req *pluginpb.CodeGeneratorRequest) {

	opts := protogen.Options{}

	plugin, err := opts.New(req)
	if err != nil {
		panic(err)
	}

	for _, file := range plugin.Files {

		log.Println(file)
	}
}
