load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = glob(["*.go"]),
    visibility = ["//visibility:public"],
    deps = [
      "//protos:go_default_proto_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//protoc-gen-go/generator:go_default_library",
        "@com_github_golang_protobuf//protoc-gen-go/descriptor:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs =  glob(["*_test.go"]),
    library = ":go_default_library",
)
