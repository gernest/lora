package imgr_test

import (
	"path"
	"path/filepath"
	"strings"

	"bitbucket.org/kardianos/osext"
	. "github.com/gernest/lora/imgr"

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
		allow = []string{".png", ".jpeg", ".PNG", ".JPG"}
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
			err = thumb.Process()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

})
