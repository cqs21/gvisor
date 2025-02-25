load("//tools:defs.bzl", "go_library", "go_test")
load("//tools/go_generics:defs.bzl", "go_template_instance")

package(licenses = ["notice"])

go_template_instance(
    name = "futex_mutex",
    out = "futex_mutex.go",
    package = "futex",
    prefix = "futexBucket",
    substrs = {
        "genericMark": "futexBucket",
    },
    template = "//pkg/sync/locking:generic_mutex",
)

go_template_instance(
    name = "atomicptr_bucket",
    out = "atomicptr_bucket_unsafe.go",
    package = "futex",
    suffix = "Bucket",
    template = "//pkg/sync/atomicptr:generic_atomicptr",
    types = {
        "Value": "bucket",
    },
)

go_template_instance(
    name = "waiter_list",
    out = "waiter_list.go",
    package = "futex",
    prefix = "waiter",
    template = "//pkg/ilist:generic_list",
    types = {
        "Element": "*Waiter",
        "Linker": "*Waiter",
    },
)

go_library(
    name = "futex",
    srcs = [
        "atomicptr_bucket_unsafe.go",
        "futex.go",
        "futex_mutex.go",
        "waiter_list.go",
    ],
    visibility = ["//pkg/sentry:internal"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/context",
        "//pkg/errors/linuxerr",
        "//pkg/hostarch",
        "//pkg/log",
        "//pkg/sentry/memmap",
        "//pkg/sync",
        "//pkg/sync/locking",
        "//pkg/usermem",
    ],
)

go_test(
    name = "futex_test",
    size = "small",
    srcs = ["futex_test.go"],
    library = ":futex",
    deps = [
        "//pkg/atomicbitops",
        "//pkg/context",
        "//pkg/errors/linuxerr",
        "//pkg/hostarch",
        "//pkg/sync",
    ],
)
