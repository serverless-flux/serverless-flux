#!/bin/sh -eu

go get github.com/vektra/mockery/.../
go get github.com/mattn/goveralls

curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.10.2
