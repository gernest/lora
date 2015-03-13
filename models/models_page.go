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

package models

import (
	"os"
	"fmt"
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
	contentPath := filepath.Join(project.ProjectPath, fmt.Sprintf("content/%s",p.Title))
		err:=os.MkdirAll(contentPath,0660)
	if err!=nil {
		logThis.Debug("Trouble %v", err)
		return err
	}
	contentFile := p.Title + ".md"
	pageFilePath := filepath.Join(contentPath, contentFile)
	content := p.Content

	buf := new(bytes.Buffer)
	fullPage := new(bytes.Buffer)
	p.Content = ""
	err = toml.NewEncoder(buf).Encode(&p)
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

func (p *Page) SaveDataFile(dest string) error {
	datatPath := filepath.Join(dest, fmt.Sprintf("data/%s",p.Title))
	err:=os.MkdirAll(datatPath,0660)
	if err!=nil {
		logThis.Debug("Trouble %v", err)
		return err
	}
	contentFile := p.Title + ".toml"
	dataFilePath := filepath.Join(datatPath, contentFile)

	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(p)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dataFilePath, buf.Bytes(), 0660)
	if err != nil {
		return err
	}
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
