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

package models_test

import (
	. "github.com/gernest/lora/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ModelsLora", func() {
	var (
		page        Page
		pages       []Page
		project     Project
		projects    []Project
		section     Section
		sections    []Section
		subsection  SubSection
		subsections []SubSection
		account     Account
		accounts    []Account
		profile     Profile
		profiles    []Profile
		lora        *Lora
	)
	BeforeEach(func() {
		lora = NewLoraObject()
	})

	It("Adds Page", func() {
		page = Page{Title: "home"}
		lora.Add(page)
		Expect(lora.HasPage).Should(BeTrue())

	})
	It("Adds Pages", func() {
		page1 := Page{Title: "Number 1"}
		page2 := Page{Title: "number two"}
		pages = []Page{page1, page2}
		lora.Add(pages)
		Expect(lora.HasPages).Should(BeTrue())
		Expect(lora.Pages[0]).Should(Equal(page1))
	})
	It("Add project", func() {
		project = Project{}
		lora.Add(project)
		Expect(lora.HasProject).Should(BeTrue())
	})
	It("Adds projects", func() {
		projects = []Project{Project{}, Project{}}
		lora.Add(projects)
		Expect(lora.HasProjects).Should(BeTrue())
	})
	It("Add section", func() {
		section = Section{}
		lora.Add(section)
		Expect(lora.HasSection).Should(BeTrue())
	})
	It("Add secions", func() {
		sections = []Section{Section{}, Section{}}
		lora.Add(sections)
		Expect(lora.HasSections).Should(BeTrue())
	})
	It("Add subasection", func() {
		subsection = SubSection{}
		lora.Add(subsection)
		Expect(lora.HasSubSection).Should(BeTrue())
	})
	It("Adds subsections", func() {
		subsections = []SubSection{SubSection{}, SubSection{}}
		lora.Add(subsections)
		Expect(lora.HasSubSections).Should(BeTrue())
	})
	It("Add account", func() {
		account = Account{}
		lora.Add(account)
		Expect(lora.HasAccount).Should(BeTrue())
	})
	It("Adds Accounts", func() {
		accounts = []Account{Account{}, Account{}}
		lora.Add(accounts)
		Expect(lora.HasAccounts).Should(BeTrue())
	})
	It("Adds profile", func() {
		profile = Profile{}
		lora.Add(profile)
		Expect(lora.HasProfile).Should(BeTrue())
	})
	It("Adds profiles", func() {
		profiles = []Profile{Profile{}, Profile{}}
		lora.Add(profiles)
		Expect(lora.HasProfiles).Should(BeTrue())
	})
})
