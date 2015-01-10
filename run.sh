#!/bin/sh

export GOPATH=$HOME/gosrc
export GOROOT=$HOME/go
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:$HOME/gosrc/bin


go get -v github.com/onsi/ginkgo/ginkgo
go get -v github.com/onsi/gomega
go get -v -t ./..

echo "Running tests"

gingko -r

echo "Running lora"

go run main.go

