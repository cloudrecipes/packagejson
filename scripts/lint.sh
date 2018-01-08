#!/bin/sh
# Source: https://github.com/mlafeldt/chef-runner/blob/master/script/lint
# Check code style and correctness.
# Usage: script/lint

set -e

golint ./...
go vet ./...