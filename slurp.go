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
	"os"
	"path/filepath"

	sh "github.com/codeskyblue/go-sh"
	"github.com/omeid/slurp"
	"github.com/omeid/slurp/stages/fs"
	"github.com/slurp-contrib/jsmin"
	"github.com/slurp-contrib/watch"
)

var (
	ProjectDir = "/home/gernest/gosrc/src/github.com/gernest/lora"
	BuildDIr   = "/home/gernest/builds/lora"
	shell      = sh.NewSession().SetDir(ProjectDir)
)

func Slurp(b *slurp.Build) {
	b.Task("clean", nil, func(c *slurp.C) error {
		return fs.Src(c,
			"static/css/*.css",
			"static/js/*.js",
		).Then(
			Remove(c),
		)
	})
	b.Task("dev", nil, func(c *slurp.C) error {
		return shell.Command("bee", "run").Run()
	})
	b.Task("test", nil, func(c *slurp.C) error {
		return shell.Command("ginkgo", "-r").Run()
	})
	b.Task("jslibs", nil, func(c *slurp.C) error {
		return fs.Src(c, "assets/js/libs/*.js").Then(
			fs.Dest(c, "static/js"),
		)
	})
	b.Task("js", nil, func(c *slurp.C) error {
		return fs.Src(c, "assets/js/frontend/*.js").Then(
			slurp.Concat(c, "lora.min.js"),
			jsmin.JSMin(c),
			fs.Dest(c, "static/js"),
		)
	})
	b.Task("css", nil, func(c *slurp.C) error {
		return fs.Src(c, "assets/css/*.css").Then(
			slurp.Concat(c, "lorastyle.css"),
			fs.Dest(c, "static/css"),
		)
	})
	b.Task("frontend", []string{"clean", "jslibs", "js", "css"}, func(c *slurp.C) error {
		return nil
	})
	b.Task("watch", []string{"frontend"}, func(c *slurp.C) error {
		j := watch.Watch(c, func(string) { b.Run(c, "js") }, "assets/js/frontend/*.js")
		css := watch.Watch(c, func(string) { b.Run(c, "css") }, "assets/css/*.css")
		b.Defer(func() {
			j.Close()
			css.Close()
		})
		b.Wait()
		return nil
	})
}

func Remove(c *slurp.C) slurp.Stage {
	return func(in <-chan slurp.File, out chan<- slurp.File) {
		for file := range in {
			err := os.Remove(file.Path)
			if err != nil {
				c.Println(err)
				return
			}
			c.Printf("Removed %s", filepath.Base(file.Path))
		}
	}
}
