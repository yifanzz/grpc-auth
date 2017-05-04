package main

import (
	"io/ioutil"
	"os"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/yifanz/grpc-auth/plugin"
	"github.com/yifanz/grpc-auth/protos/go_default_proto_library"
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
	for _, f := range request.ProtoFile {
		generate(f)
	}
}

func generate(file *descriptor.FileDescriptorProto) {
	for _, svc := range file.Service {
		for _, m := range svc.Method {
			if m.Options == nil || !proto.HasExtension(m.Options, options.E_AuthOptions) {
				continue
			}
			rawOptions, _ := proto.GetExtension(m.Options, options.E_AuthOptions)
			authOptions := rawOptions.(*options.AuthMethodOptions)

			fmt.Printf("got method %s that has auth option %+v", *m.Name, authOptions.Simple.Path)
		}
	}
}
