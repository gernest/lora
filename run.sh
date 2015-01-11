#!/bin/sh
# Copyright 2015 Geofrey Ernest a.k.a gernest, All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"): you may
#  not use this file except in compliance with the License. You may obtain
#  a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.

# This scripts helps to run tests and lora without building
#
#  ****THIS NEEDS A POSTGRESQL DATABASE INORDER TO WORK*******
#
# Please make sure you have a working postgres database and edit the file conf/app.conf
# to reflect your database settings
export GOPATH=$HOME/gosrc
export GOROOT=$HOME/go
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:$HOME/gosrc/bin


go get -v github.com/onsi/ginkgo/ginkgo
go get -v github.com/onsi/gomega
go get -v -t ./...

echo "Running tests....."

ginkgo -r

echo "Running lora....."

go run main.go

