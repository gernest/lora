package models_test

import (
	. "github.com/gernest/lora/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ModelsSection", func() {
	var (
		section    *Section
		subsection *SubSection
		body       string
	)
	Describe("Section", func() {
		BeforeEach(func() {
			section = new(Section)
			body = "# mambo"
			section.Body = body
		})
		It("", func() {
			section.Sanitize()
			x := SanitizeTestHelper(body)
			Expect(section.BodyHtml).Should(Equal(x))
		})
	})
	Describe("SuSection", func() {
		BeforeEach(func() {
			subsection = new(SubSection)
			body = "# pouwa"
			subsection.Body = body
		})
		It("", func() {
			subsection.Sanitize()
			x := SanitizeTestHelper(body)
			Expect(subsection.BodyHtml).Should(Equal(x))
		})
	})
})
