load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = glob(["*.go"]),
    deps = [
      "//plugin:go_default_library",
      "@com_github_golang_protobuf//proto:go_default_library",
      "@com_github_golang_protobuf//protoc-gen-go/plugin:go_default_library",
      "@com_github_golang_protobuf//protoc-gen-go/descriptor:go_default_library",
    ],
)

go_binary(
    name = "protoc-gen-grpc-auth",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
