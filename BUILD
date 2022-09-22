# Copyright 2014 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
    out = "protoc-gen-bq-schema",
    deps = [
        "//internal/converter:bq_schema_go_lib",
    ] + COMP_DEPS,
)
