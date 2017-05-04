load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_proto_library(
  name = "go_default_proto_library",
  protos = glob(["*.proto"]),
  langs = [
    "@org_pubref_rules_protobuf//go",
    "//:grpc-auth",
  ],
  verbose = 0,
)

go_test(
  name = "default_go_test",
  srcs = glob(["*_test.go"]),
  library = ":go_default_proto_library"
)
