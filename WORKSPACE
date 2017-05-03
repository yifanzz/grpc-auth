git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/improbable-io/rules_go.git",
    commit = "b05526382b891b6287471111b020cea86b689921",
)
load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository", "go_repository")

go_repositories()

git_repository(
  name = "org_pubref_rules_protobuf",
  remote = "https://github.com/improbable-io/rules_protobuf",
  commit = "46f6aa2b3a966dbd40b3ecacb6278bb080248184"
)

load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")

go_proto_repositories(
  excludes = [
    "com_github_golang_protobuf",
    "org_golang_google_grpc",
    "org_golang_x_net",
  ]
)

new_go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    remote = "https://github.com/improbable-io/protobuf",
    vcs = "git",
    commit = "665f23bbc8394d7f2f01f1f107310e33f542a224",
)

new_go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    commit = "f9114dace7bd920b32f943b3c73fafbcbab2bf31",
)

new_go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    remote = "https://github.com/improbable-io/grpc-go",
    vcs = "git",
    commit = "75991db8728087a7c3ea2a730419cb876740d71e",
)

new_go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    remote = "https://github.com/golang/net",
    vcs = "git",
    commit = "da118f7b8e5954f39d0d2130ab35d4bf0e3cb344",
)

new_go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    remote = "https://github.com/golang/text",
    vcs = "git",
    commit = "44f4f658a783b0cee41fe0a23b8fc91d9c120558",

    build_file_name = "BUILD.bazel",
)

new_go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    remote = "https://github.com/google/go-genproto",
    vcs = "git",
    commit = "b3e7c2fb04031add52c4817f53f43757ccbf9c18",
)
