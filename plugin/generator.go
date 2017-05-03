package plugin

import (
	// "fmt"
	//
	// "github.com/gogo/protobuf/gogoproto"
	// "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

func init() {
	generator.RegisterPlugin(NewSimpleAuthPlugin())
}

type simpleAuthPlugin struct {
	*generator.Generator
	generator.PluginImports
}

func NewSimpleAuthPlugin() *simpleAuthPlugin {
	return &simpleAuthPlugin{}
}

func (p *simpleAuthPlugin) Name() string {
	return "auth"
}

func (p *simpleAuthPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *simpleAuthPlugin) Generate(file *generator.FileDescriptor) {
	println("generating some stuf")
	for _, msg := range file.Messages() {
		println(">>>> generating for ", *msg.Name)
	}
	p.P("//test string")
}

func (p *simpleAuthPlugin) GenerateImports(file *generator.FileDescriptor) {
}
