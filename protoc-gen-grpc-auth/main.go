package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/yifanz/grpc-auth/plugin"
	"strings"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	request := &plugin_go.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, request); err != nil {
		return
	}


	fmt.Fprintf(os.Stderr, ">>>>> generating for file %+v\n", request.FileToGenerate)

	var files []*plugin_go.CodeGeneratorResponse_File
	for _, f := range request.GetProtoFile() {
		generated, content, _ := plugin.Generate(f)
		if generated {
			files = append(files, &plugin_go.CodeGeneratorResponse_File{
				Name: proto.String(strings.Replace(f.GetName(), ".proto", ".pb.auth.go", -1)),
				Content: proto.String(content),
			})
		}
	}
	emitFiles(files)
}

func emitFiles(out []*plugin_go.CodeGeneratorResponse_File) {
	emitResp(&plugin_go.CodeGeneratorResponse{File: out})
}

func emitError(err error) {
	emitResp(&plugin_go.CodeGeneratorResponse{Error: proto.String(err.Error())})
}

func emitResp(resp *plugin_go.CodeGeneratorResponse) {
	buf, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("failed to generate: %v", err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatal(err)
	}
}
