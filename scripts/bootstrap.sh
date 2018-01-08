#!/bin/sh
# Source: https://github.com/mlafeldt/chef-runner/blob/master/script/bootstrap
# Install Go dependencies.
# Usage: script/bootstrap

set -e

go get -d -t ./...
go get -v \
  github.com/golang/lint/golint \
  golang.org/x/tools/cmd/cover \
  github.com/mattn/goveralls

