package server_template

import (
	"bytes"
	"github.com/bradleyaus/mockymcmockface/pkg/parser"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"log"
	"text/template"
)

type parameters struct {
	Services []Service
	Methods  []Method
	Imports []Import
}

type Options struct {
	TemplateName string
	TemplatePath string
}


type Service struct {
	Name string
	Import string
}

//TODO: Support all method types, eg stream
type Method struct {
	MethodName  string
	ServiceName string

	InputImport string
	OutputImport string

	InputType   string
	OutputType  string
}

func GenerateFromTemplate(services []parser.Service, deps parser.ProtoToPackageImport, opt *Options) ([]byte, error) {


	imp, aliasMap := createImports(deps)
	parameters := flatten(services, aliasMap)
	parameters.Imports = imp

	tmplString, err := getTemplateString(opt.TemplatePath)
	if err != nil {
		log.Println("failed when loading template", err)
		return nil, err
	}

	tmpl := template.New(opt.TemplateName)
	tmpl, err = tmpl.Parse(tmplString)
	if err != nil {
		log.Println("failed when parsing template", err)
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, parameters)
	if err != nil {
		log.Println("failed when executing template", err)
		return nil, err
	}

	b := buf.Bytes()
	formatted, err := imports.Process("", b, nil)
	if err != nil {
		log.Println("failed when formatting", err)
		log.Print(string(b))
		return nil, err
	}

	return formatted, nil
}

func getTemplateString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

