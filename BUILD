load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
load("@org_pubref_rules_protobuf//protobuf:rules.bzl", "proto_language")

go_prefix("github.com/yifanz/grpc-auth", ".")

proto_language(
    name = "grpc-auth",
    grpc_file_extensions = [".pb.auth.go"],
    grpc_inputs = [
        "//protos:default_protos",
        "@com_github_google_protobuf//:well_known_protos",
    ],
    grpc_imports = [
        "external/com_github_google_protobuf/src/",
    ],
    grpc_options = ["generate_empty=true"],
    grpc_plugin = "//protoc-gen-grpc-auth",
    grpc_plugin_name = "grpc-auth",
    importmap = {
      "google/protobuf/descriptor.proto": "github.com/golang/protobuf/protoc-gen-go/descriptor",
    },
    prefix = ":go_prefix",
    supports_grpc = True,
    supports_pb = True,
    visibility = ["//visibility:public"],
)
