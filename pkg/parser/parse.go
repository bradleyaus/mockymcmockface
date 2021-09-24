package parser

import "google.golang.org/protobuf/compiler/protogen"

type Service struct {
	Name       string
	Proto      string
	Package    string
	ImportPath string
	Methods    []Method
}

//TODO: Support all method types, eg stream
type Method struct {
	Name            string
	Proto           string
	InputVariables  []Variable
	OutputVariables []Variable
}

type Variable struct {
	Name  string
	Proto string
}

func ParseServices(plugin *protogen.Plugin) []Service {
	var svcs []Service

	for _, file := range plugin.Files {

		for _, s := range file.Services {
			svcs = append(svcs, extractService(file, s))
		}
	}

	return svcs
}

func extractService(file *protogen.File, ps *protogen.Service) (s Service) {
	s = Service{
		Name:       ps.GoName,
		Proto:      ps.Location.SourceFile,
		Package:    string(file.GoPackageName),
		ImportPath: file.GoImportPath.String(),
		Methods:    extractMethods(ps),
	}

	return s
}

func extractMethods(ps *protogen.Service) (methods []Method) {
	for _, method := range ps.Methods {
		methods = append(methods, Method{
			Name:            method.GoName,
			Proto:           method.Location.SourceFile,
			InputVariables:  extractVariables(method.Input),
			OutputVariables: extractVariables(method.Output),
		})
	}

	return methods
}

func extractVariables(m *protogen.Message) (variables []Variable) {
	for _, field := range m.Fields {
		variables = append(variables, Variable{
			Name:  m.GoIdent.GoName,
			Proto: field.Location.SourceFile,
		})
	}
	return variables
}
