#!/bin/sh

export GOPATH=$HOME/gosrc
export GOROOT=$HOME/go
export PATH=$PATH:$HOME/go/bin

go get -v ./..

go run main.go

