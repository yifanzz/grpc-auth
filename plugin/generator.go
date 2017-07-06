package plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/yifanz/grpc-auth/protos/go_default_proto_library"

	"bytes"
)

func Generate(file *descriptor.FileDescriptorProto) (bool, string, error) {
	out := bytes.NewBuffer(nil)
	fileGenerated := false
	for _, svc := range file.GetService() {
		for _, m := range svc.GetMethod() {
			if m.Options == nil || !proto.HasExtension(m.Options, options.E_AuthOptions) {
				continue
			}
			rawOptions, _ := proto.GetExtension(m.Options, options.E_AuthOptions)
			authOptions := rawOptions.(*options.AuthMethodOptions)

			if authOptions.Simple != nil {
				fileGenerated = true
				err := simpleAuthTemplate.Execute(out, &params{})
				if err != nil {
					return false, "", err
				}
			}
		}
	}
	return fileGenerated, out.String(), nil
}

type params struct {

}