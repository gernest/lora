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

package imgr_test

import (
	"path"
	"path/filepath"
	"strings"

	. "github.com/gernest/lora/imgr"
	"github.com/kardianos/osext"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Imgr", func() {
	var (
		base    string
		manager *ImageManager
		err     error
		src     string
		dst     string
		allow   []string
	)
	BeforeEach(func() {
		basePath, _ := osext.ExecutableFolder()
		base = filepath.Join(path.Dir(strings.TrimSuffix(basePath, "/")), "fixtures")
		src = filepath.Join(base, "imgr/src")
		dst = filepath.Join(base, "imgr/dst")
		allow = []string{".png", ".jpeg", ".PNG", ".JPG", ".jpg"}
	})

	Describe("ImageManager", func() {
		BeforeEach(func() {
			manager = NewImageManager(src, dst, allow)

		})
		It("loads", func() {
			err = manager.LoadFromSource()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(manager.Images).ShouldNot(BeEmpty())
		})

	})
	Describe("ThumbailManager", func() {
		var thumb *Thumbnails
		BeforeEach(func() {
			thumb = &Thumbnails{}
			thumb.Source = src
			thumb.Destinalion = dst
			thumb.AllowedExt = allow
		})
		It("Loads", func() {
			err = thumb.LoadFromSource()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Process", func() {
			_ = thumb.LoadFromSource()
			err = thumb.Process(0, 0)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Create a single thumnail", func() {
			source := filepath.Join(thumb.Source, "tusha.png")
			err = thumb.CreateThumbnail(source, thumb.Destinalion, 200, 200)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("SHould make tiny stuffs", func() {
			d := thumb.Destinalion
			thumb.Destinalion = filepath.Join(d, "tiny")
			_ = thumb.LoadFromSource()
			err = thumb.Process(100, 100)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

})
