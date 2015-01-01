package models_test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"bitbucket.org/kardianos/osext"
	. "github.com/gernest/lora/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project", func() {
	var (
		base           string
		err            error
		project        Project
		basePath       string
		currentProject *Project
		baseProject    Project
	)
	baseProject = Project{
		Id:           1,
		Title:        "my new tushabe site",
		Name:         "pasiansi",
		Theme:        "loraina",
		PublishDir:   "www",
		BaseUrl:      "http://yourSiteHere",
		LanguageCode: "en-us",
		Pages: []Page{
			{Id: 1, Title: "home", Content: "## hello home", Slug: "slug", Draft: false},
			{Id: 2, Title: "about", Content: "## hello about", Slug: "slug", Draft: false},
			{Id: 3, Title: "products", Content: "## hello products", Slug: "slug", Draft: false},
			{Id: 4, Title: "contact", Content: "## hello contact", Slug: "slug", Draft: false},
		},
	}
	basePath, _ = osext.ExecutableFolder()
	base = filepath.Join(path.Dir(strings.TrimSuffix(basePath, "/")), "fixtures")
	Describe("InitializeProject", func() {
		AfterEach(func() {
			clearAll(filepath.Join(base, "projects"))
		})
		PContext(" Given the base path ", func() {

			It("Shold populate with initial values", func() {
				p, _:= NewLoraProject(base, "mchele", "", "")
				Expect(base).Should(Equal(p.BaseDir))
			})

		})
		Context("Without base path", func() {

			It("should have default values ", func() {
				p := new(Project)
				err = p.Initialize(project.BaseDir, "unga", "", "")
				Expect(err).Should(HaveOccurred())
			})
		})

	})
	Describe("GenerateScaffold", func() {
		BeforeEach(func() {
			currentProject = new(Project)
			_ = currentProject.Initialize(base, "kilimahewa", "", "")

		})
		AfterEach(func() {
			_ = currentProject.Clean()
		})

		It("should generate project", func() {
			err = currentProject.GenScaffold()
			projectPath := base + "/projects/" + "kilimahewa"
			file, _ := os.Stat(currentProject.ProjectPath)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(currentProject.ProjectPath).Should(Equal(projectPath))
			Expect(file.IsDir()).Should(BeTrue())
		})
	})

	Describe("Clean", func() {
		BeforeEach(func() {
			currentProject = new(Project)
			_ = currentProject.Initialize(base, "bigbite", "", "")
			_ = currentProject.GenScaffold()

		})

		It("should remove generated files", func() {
			err = currentProject.Clean()
			file, _ := os.Stat(currentProject.ProjectPath)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(file).Should(BeNil())
		})

	})
	Describe("LoadConfigFile", func() {
		BeforeEach(func() {
			currentProject = new(Project)
			_ = currentProject.Initialize(base, "pasiansi", "", "")
			_ = currentProject.GenScaffold()
		})
		AfterEach(func() {
			_ = currentProject.Clean()
		})
		It("Loads config file", func() {
			err = currentProject.LoadConfigFile()

			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should have correct confg values", func() {
			_ = currentProject.LoadConfigFile()

			Expect(currentProject.Name).Should(Equal(baseProject.Name))
			Expect(currentProject.Title).Should(Equal(baseProject.Title))
			Expect(currentProject.Id).Should(Equal(baseProject.Id))
			Expect(currentProject.PublishDir).Should(Equal(baseProject.PublishDir))
			Expect(currentProject.BaseUrl).Should(Equal(baseProject.BaseUrl))
		})
	})
	Describe("SaveConfigFile", func() {
		BeforeEach(func() {
			_ = currentProject.Initialize(base, "pasiansi", "", "")
			_ = currentProject.GenScaffold()
			_ = currentProject.LoadConfigFile()
		})
		AfterEach(func() {
			_ = currentProject.Clean()
		})
		It("Should save", func() {
			err = currentProject.SaveConfigFile()

			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should update the values", func() {
			currentProject.Title = "Fuck ISIS"
			_ = currentProject.SaveConfigFile()
			_ = currentProject.LoadConfigFile()

			Expect(currentProject.Title).Should(Equal("Fuck ISIS"))
		})
	})

	Describe("Page", func() {
		BeforeEach(func() {
			_ = currentProject.Initialize(base, "pasiansi", "", "")
			_ = currentProject.GenScaffold()
			_ = currentProject.LoadConfigFile()
		})
		AfterEach(func() {
			_ = currentProject.Clean()
		})
		It("Should generate a new page", func() {
			p := currentProject.Pages[0]

			Expect(p.Generate(currentProject)).Should(Succeed())
			Expect(p.ContentPath).ShouldNot(BeEmpty())
		})
	})

})

func clearAll(s string) {
	fmt.Printf("cleaning %s \n", s)
	_ = os.RemoveAll(s)
}