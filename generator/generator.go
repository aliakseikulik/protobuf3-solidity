package generator

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

// SolidityVersionString is the Solidity version specifier.
const SolidityVersionString = ">=0.6.0 <8.0.0"

// Generator generates Solidity code from .proto files.
type Generator struct {
	request *pluginpb.CodeGeneratorRequest
}

// New initializes a new Generator.
func New(request *pluginpb.CodeGeneratorRequest) *Generator {
	g := new(Generator)
	g.request = request
	return g
}

// Generate generates Solidity code from the requested .proto files.
func (g *Generator) Generate() (*pluginpb.CodeGeneratorResponse, error) {
	println(proto.MarshalTextString(g.request))

	// println(g.request.GetFileToGenerate()[0])

	return nil, nil
}
