load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "openscenario-1.3.1",
    srcs = ["OpenSCENARIO_1.3.1.xsd"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "openscenario",
    srcs = ["openscenario_schema.go"],
    importpath = "github.com/gerzin/xoscs/backend/backend/pkg/xosc/schemas",
    visibility = ["//visibility:public"],
)
