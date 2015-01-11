package models_test

import (
	. "github.com/gernest/lora/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ModelsPage", func() {
	var (
		page    *Page
		content string
		err     error
	)
	BeforeEach(func() {
		page = new(Page)
		content = "# mambo"
	})
	It("Sanitize", func() {
		page.Content = content
		page.Sanitize()
		x := SanitizeTestHelper(content)
		Expect(page.ContentHtml).Should(Equal(x))
	})
	It("", func() {
		p := new(Project)
		err = page.Generate(p)
		Expect(err).Should(HaveOccurred())
	})
})
