load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "context",
    testonly = 1,
    srcs = ["context.go"],
    visibility = [
        "//visibility:public",
    ],
    deps = [
        "//pkg/buffer",
        "//pkg/tcpip",
        "//pkg/tcpip/checker",
        "//pkg/tcpip/header",
        "//pkg/tcpip/link/channel",
        "//pkg/tcpip/link/sniffer",
        "//pkg/tcpip/network/ipv4",
        "//pkg/tcpip/network/ipv6",
        "//pkg/tcpip/seqnum",
        "//pkg/tcpip/stack",
        "//pkg/tcpip/transport/tcp",
        "//pkg/waiter",
    ],
)
