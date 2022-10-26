#!/bin/bash

set -ex

export CGO_ENABLED=0 GOOS=linux GOARCH=amd64

cd cmd/fy
go build -tags embed -v -o ../../bin/fy-embed .
go build -tags external -v -o ../../bin/fy-external .
cd ../..

mkdir -p "_build/sbin"
mkdir -p "_build/etc/systemd/system/"
cp hack/service/fy.service "_build/etc/systemd/system/fy.service"
