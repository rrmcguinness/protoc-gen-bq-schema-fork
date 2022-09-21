load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc//doc:defs.bzl", "doc_markdown_compile")

package(default_visibility = ["//:__subpackages__"])

proto_library(
    name = "bq_schema_proto_lib",
    srcs = [
        "bq_field.proto",
        "bq_table.proto",
    ],
    deps = [
        "@com_google_protobuf//:descriptor_proto",
    ],
)

doc_markdown_compile(
    name = "docs",
    protos = [
        ":bq_schema_proto_lib",
    ],
)
