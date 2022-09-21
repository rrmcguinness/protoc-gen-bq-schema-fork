load("@bazel_gazelle//:def.bzl", "gazelle")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")
load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("//:conf/deps.bzl", "COMP_DEPS")

package(default_visibility = ["//:__subpackages__"])

buildifier(
    name = "buildifier",
)

gazelle(
    name = "gazelle",
    prefix = "google_retail_platform",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=conf/go_deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_binary(
    name = "protoc-gen-bq-schema",
    srcs = [
        "main.go",
    ],
    out = "bq-schema-generator",
    deps = [
        "//internal/converter:bq_schema_go_lib",
    ] + COMP_DEPS,
)
