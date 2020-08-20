load("//tools:defs.bzl", "go_binary", "go_library")

package(licenses = ["notice"])

go_library(
    name = "test",
    srcs = ["test_unsafe.go"],
)

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [":test"],
)
