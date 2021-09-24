package main

import (
	"github.com/bradleyaus/mockymcmockface/pkg/parser"
	"github.com/bradleyaus/mockymcmockface/pkg/server_template"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	// This is a compiler that is called by protoc, which means we have to compile this into a binary and run.
	// It passes all that is required for compilation through stdin
	// The problem with that is it's difficult to quickly iterate on a solution
	// So to make development easier/allow debugging, we can save the input from stdin to a file and then use that in a examples
	ioutil.WriteFile("input.dat", input, os.ModePerm)

	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	params := make(map[string]string)
	for _, param := range strings.Split(req.GetParameter(), ",") {
		split := strings.Split(param, "=")
		params[split[0]] = split[1]
	}

	templatePath := "server.tmpl"
	if params["template"] != "" {
		templatePath = params["template"]
	}

	plugin := generate(&req, &server_template.Options{
		TemplateName: "server.tmpl",
		TemplatePath: templatePath,
	})

	resp := plugin.Response()
	var supportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	resp.SupportedFeatures = &supportedFeatures
	out, err := proto.Marshal(resp)
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
