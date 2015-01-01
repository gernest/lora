package models

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

// Sanitize This prepares the ContentHtml field for rendering safe html
func (p *Page) Sanitize() {
	p.ContentHtml = template.HTML(sanitizeHTMLField(p.Content))
}

func (p *Page) LastUpdate() string {
	return p.UpdatedAt.Format(time.RFC822)
}

// Generate creates a markdown file for the given page
func (p *Page) Generate(project *Project) error {
	var fm string
	contentPath := filepath.Join(project.ProjectPath, "content")
	contentFile := p.Title + "/" + p.Title + ".md"
	pageFilePath := filepath.Join(contentPath, contentFile)
	content := p.Content

	buf := new(bytes.Buffer)
	fullPage := new(bytes.Buffer)
	p.Content = ""
	err := toml.NewEncoder(buf).Encode(&p)
	if err != nil {
		clean := project.Clean()
		if clean != nil {
			return clean
		}
		return err
	}
	fm = frontmatter(buf.String())

	fullPage.WriteString(fm)
	fullPage.WriteString(content)

	err = ioutil.WriteFile(pageFilePath, fullPage.Bytes(), 0660)
	if err != nil {
		clean := project.Clean()
		if clean != nil {
			return clean
		}
		return err
	}
	p.ContentPath = pageFilePath
	return nil

}

func frontmatter(s string) string {
	separator := "+++"
	buf := new(bytes.Buffer)
	buf.WriteString(separator)
	buf.WriteString("\n")
	buf.WriteString(s)
	buf.WriteString("\n")
	buf.WriteString(separator)
	buf.WriteString("\n\n")
	return buf.String()
}
