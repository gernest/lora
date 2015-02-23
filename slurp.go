// +build slurp
// Copyright 2015 Geofrey Ernest a.k.a gernest, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"

	sh "github.com/codeskyblue/go-sh"
	"github.com/omeid/slurp"
	"github.com/omeid/slurp/stages/fs"
)

var BuildDIr string
var ProjectDir string

func init() {
	ProjectDir = "/home/gernest/gosrc/src/github.com/gernest/lora"
	BuildDIr = "/home/gernest/builds/lora"
}
func Slurp(b *slurp.Build) {
	b.Task("clean", nil, func(c *slurp.C) error {
		sess := NewSession(ProjectDir)
		return sess.Command("bundle", "exec", "compass", "clean").Run()
	})

	b.Task("compass", []string{"clean"}, func(c *slurp.C) error {
		compass := NewSession(ProjectDir)
		return compass.Command("bundle", "exec", "compass", "watch").Run()
	})
	b.Task("dev", nil, func(c *slurp.C) error {
		bee := NewSession(ProjectDir)
		return bee.Command("bee", "run").Run()
	})
	b.Task("test", nil, func(c *slurp.C) error {
		test := NewSession(ProjectDir)
		return test.Command("ginkgo", "-r").Run()
	})
	b.Task("views", nil, func(c *slurp.C) error {
		return fs.Src(c,
			fmt.Sprintf("%s/views/*", ProjectDir),
		).Then(
			fs.Dest(c, fmt.Sprintf("%s/views", BuildDIr)),
		)
	})
}

func NewSession(path string) *sh.Session {
	return sh.NewSession().SetDir(path)
}
