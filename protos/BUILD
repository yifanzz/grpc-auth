load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_library")

filegroup(
    name = "default_protos",
    srcs = [
        "options.proto",
    ],
    visibility = ["//visibility:public"],
)

go_proto_library(
  name = "go_default_proto_library",
  protos = [ ":default_protos" ],
  langs = [
    "@org_pubref_rules_protobuf//go",
  ],
  inputs = [
    "@com_github_google_protobuf//:well_known_protos",
  ],
  imports = [
    "external/com_github_google_protobuf/src/",
  ],
  importmap = {
    "google/protobuf/descriptor.proto": "github.com/golang/protobuf/protoc-gen-go/descriptor",
  },
  deps = [
    "@com_github_golang_protobuf//protoc-gen-go/descriptor:go_default_library",
  ],
  visibility = ["//visibility:public"],
  verbose = 2, # 0=no output, 1=show protoc command, 2+ more...
  prefix = "//:go_prefix"
)
