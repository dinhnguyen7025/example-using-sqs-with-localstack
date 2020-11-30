#!/bin/sh
set -eux

cd "`dirname $0`"

export GOOS=linux
export GOARCH=amd64

go mod tidy -v

main_app=./cmd/custom
dest=bin/custom

./cmd/build ${main_app} ${dest}
