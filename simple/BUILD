load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_proto_library(
  name = "go_default_proto_library",
  protos = glob([ "*.proto" ]),
  langs = [
    "@org_pubref_rules_protobuf//go",
    "//:grpc-auth",
  ],
  deps = [
          "//protos:go_default_proto_library",
  ],
  importmap = {
          "grpc/auth/options.proto": "github.com/yifanz/grpc-auth/protos/go_default_proto_library",
  },
  with_grpc = True,
)

go_test(
    name = "go_default_test",
    srcs =  glob([ "*_test.go" ]),
    library = ":go_default_proto_library",
)