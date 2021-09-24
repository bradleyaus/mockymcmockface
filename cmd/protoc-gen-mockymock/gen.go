package main

import (
	"github.com/bradleyaus/mockymcmockface/pkg/parser"
	"github.com/bradleyaus/mockymcmockface/pkg/server_template"
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
	// So to make development easier/allow debugging, we can save the input from stdin to a file and then use that in a examples
	os.MkdirAll("examples/protoc-gen-data/", os.ModePerm)
	ioutil.WriteFile("examples/protoc-gen-data/input.dat", input, os.ModePerm)

	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)


	plugin := generate(&req, &server_template.Options{
		TemplateName: "server.tmpl",
		TemplatePath: "server.tmpl",
	})

	out, err := proto.Marshal(plugin.Response())
	if err != nil {
		log.Fatalln("failed when marshaling response")
	}

	os.Stdout.Write(out)
}

func generate(req *pluginpb.CodeGeneratorRequest, options *server_template.Options) (*protogen.Plugin) {

	opts := protogen.Options{}

	plugin, err := opts.New(req)
	if err != nil {
		panic(err)
	}

	services := parser.ParseServices(plugin)
	deps := parser.GenerateDependencies(plugin)

	b, err := server_template.GenerateFromTemplate(services, deps, options)
	if err != nil {
		log.Fatalln("failed when generating server")
	}

	file := plugin.NewGeneratedFile("server.go", ".")
	_, err = file.Write(b)
	if err != nil {
		log.Fatalln("failed when writing output")
	}

	return plugin
}
