load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "filter",
    srcs = [
        "config.go",
        "config_amd64.go",
        "config_arm64.go",
        "extra_filters.go",
        "extra_filters_msan.go",
        "extra_filters_race.go",
        "extra_filters_race_amd64.go",
        "extra_filters_race_arm64.go",
        "filter.go",
    ],
    visibility = [
        "//runsc:__subpackages__",
    ],
    deps = [
        "//pkg/abi/linux",
        "//pkg/flipcall",
        "//pkg/log",
        "//pkg/seccomp",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)
