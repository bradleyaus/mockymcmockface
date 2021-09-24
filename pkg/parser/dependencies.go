package parser

import "google.golang.org/protobuf/compiler/protogen"

// Map proto file (eg examples.proto) to the package
type ProtoToPackageImport = map[string]string

func GenerateDependencies(plugin *protogen.Plugin) ProtoToPackageImport {

	m := ProtoToPackageImport{}
	for _, file := range plugin.Files {
		m[*file.Proto.Name] = file.GoImportPath.String()
	}

	return m
}

