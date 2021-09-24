package server_template

import parser "github.com/bradleyaus/mockymcmockface/pkg/parser"

func flatten(services []parser.Service, deps ProtoToAliasMap) (p *parameters) {

	p = &parameters{}

	for _, s := range services {
		serviceMethods := flattenMethods(s, deps)
		p.Methods = append(p.Methods, serviceMethods...)
		p.Services = append(p.Services, flattenService(s, deps))
	}

	return p
}

func flattenService(inService parser.Service, deps ProtoToAliasMap) (outService Service) {

	return Service{
		Name:   inService.Name,
		Import: deps[inService.Proto],
	}

	return outService
}

func flattenMethods(inService parser.Service, deps ProtoToAliasMap) (outMethods []Method) {

	for _, method := range inService.Methods {
		outMethods = append(outMethods, Method{
			MethodName:   method.Name,
			ServiceName:  inService.Name,
			InputImport:  deps[method.InputVariables[0].Proto],
			OutputImport: deps[method.OutputVariables[0].Proto],
			//TODO: DO i need to support multiple vars?
			InputType:  method.InputVariables[0].Name,
			OutputType: method.OutputVariables[0].Name,
		})
	}

	return outMethods
}

