package models_test

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"bitbucket.org/kardianos/osext"
	. "github.com/gernest/lora/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Profile", func() {
	var (
		prof *Profile
		err  error
		base string
	)
	BeforeEach(func() {
		basePath, _ := osext.ExecutableFolder()
		base = filepath.Join(path.Dir(strings.TrimSuffix(basePath, "/")), "fixtures")
		prof = new(Profile)
	})
	AfterEach(func() {
		cleanUp()
	})
	It("tick", func() {
		err = prof.GenerateIdenticon(base, "kilimahewa")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(prof.Photo).ShouldNot(BeEmpty())
	})
	It("photo", func() {
		_ = prof.GenerateIdenticon(base, "kilimahewa")
		Expect(prof.Photo).ShouldNot(BeEmpty())
	})
	It("crash", func() {
		err = prof.GenerateIdenticon("/", "kilimahewa")
		Expect(err).Should(HaveOccurred())
	})
	Describe("Defaults", func() {
		var (
			dest string
		)
		BeforeEach(func() {
			b, _ := osext.ExecutableFolder()
			dest = filepath.Join(b, "static")
			prof = new(Profile)
		})

		AfterEach(func() {
			_ = os.RemoveAll(dest)
		})
		It("emptyh", func() {
			err = prof.GenerateIdenticon("", "kilimahewa")
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should tick", func() {
			_ = prof.GenerateIdenticon("", "kilimahewa")
			_, err = os.Stat(dest)
			Expect(err).ShouldNot(HaveOccurred())

		})
	})

})

func cleanUp() {
	basePath, _ := osext.ExecutableFolder()
	base := filepath.Join(path.Dir(strings.TrimSuffix(basePath, "/")), "fixtures")
	dstPath := filepath.Join(base, "imgr/dst")
	removeSafely(dstPath)
}

func removeSafely(p string) {
	info, err := os.Stat(p)
	if err != nil {
		panic(err)
	}
	err = os.RemoveAll(p)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(p, info.Mode())
	if err != nil {
		panic(err)
	}
}
