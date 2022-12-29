#!/bin/bash

set -e

# install husky package
go install github.com/go-courier/husky/cmd/husky@latest

# init husky
husky init