package server_template

import (
	"fmt"
	"github.com/bradleyaus/mockymcmockface/pkg/parser"
)

type Import struct {
	Alias string
	Package string
}

type ProtoToAliasMap = map[string]string

func createImports(deps parser.ProtoToPackageImport) ([]Import, ProtoToAliasMap) {

	var imports []Import
	resultMap := ProtoToAliasMap{}
	counter := 1
	for key, val := range deps {
		alias := fmt.Sprintf("pb%v", counter)
		imports = append(imports, Import{
			Alias:   fmt.Sprintf("pb%v", counter),
			Package: val,
		})

		resultMap[key] = alias
		counter++
	}

	return imports, resultMap
}
