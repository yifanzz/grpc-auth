load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = glob(["*.go"]),
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_gogo_protobuf//proto:go_default_library",
        "@com_github_gogo_protobuf//gogoproto:go_default_library",
        "@com_github_gogo_protobuf//protoc-gen-gogo/descriptor:go_default_library",
        "@com_github_gogo_protobuf//protoc-gen-gogo/generator:go_default_library",
    ],
)
