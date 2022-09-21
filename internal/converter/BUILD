load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("//:conf/deps.bzl", "COMP_DEPS")

go_proto_library(
    name = "bq_schema_go_proto_lib",
    importpath = "github.com/GoogleCloudPlatform/protoc-gen-bq-schema/protos",
    proto = "//api:bq_schema_proto_lib",
    visibility = ["//:__subpackages__"],
)

go_library(
    name = "bq_schema_go_lib",
    srcs = [
        "comments.go",
        "convert.go",
        "protopackage.go",
    ],
    importpath = "github.com/GoogleCloudPlatform/protoc-gen-bq-schema/converter",
    visibility = ["//:__subpackages__"],
    deps = [
        ":bq_schema_go_proto_lib",
    ] + COMP_DEPS,
)

go_test(
    name = "test",
    srcs = glob(["*_test.go"]),
    embed = [":bq_schema_go_lib"],
    deps = COMP_DEPS,
)
