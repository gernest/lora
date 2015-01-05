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
	"errors"
	"os"

	"path/filepath"

	"bitbucket.org/kardianos/osext"

	"bytes"
	"io/ioutil"
	lorem "github.com/drhodes/golorem"
	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego"
	sh "github.com/codeskyblue/go-sh"
	cp "github.com/gernest/lora/utilities/copy"
)

// GenScaffold copies  a directory from thr templates folder into the projects folder.
// The template name should be provided by the user, if not default template is used
func (p *Project) GenScaffold() error {
	err := cp.CopyDir(p.TemplatePath, p.ProjectPath)
	if err != nil {
		return err
	}
	return nil
}

// InitiDir initializes the project directory by performing git stuffs
func (p *Project) InitDir() error {
	sess := sh.NewSession().SetDir(p.ProjectPath)

	// Initi a git repo
	err := sess.Command("git", "init").Run()
	if err != nil {
		clean := p.Clean()
		if clean != nil {
			return clean
		}
		return err
	}

	// Adding all files to git
	err = sess.Command("git", "add", ".").Run()
	if err != nil {
		clean := p.Clean()
		if clean != nil {
			return clean
		}
		return err
	}

	// Making a first commit
	err = sess.Command("git", "commit", "-m", "Initial Commit").Run()
	if err != nil {
		clean := p.Clean()
		if clean != nil {
			return clean
		}
		return err
	}
	return nil
}

// LoadConfig  reads the configuration file found in a project path. The file is expected
// to be of toml format, it unamrshall the values into the current project object.
func (p *Project) LoadConfigFile() error {
	configPath := filepath.Join(p.ProjectPath, "config.toml")
	_, err := toml.DecodeFile(configPath, p)
	if err != nil {
		clean := p.Clean()
		if clean != nil {
			return clean
		}
		return err
	}
	return nil
}

// GenContent  generates frontmatter and contents of the pages in Markdown files, the fuction
// can only be called after the project configurations have been loaded
func (p *Project) GenContent() error {
	for _, v := range p.Pages {
		err := v.Generate(p)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveConfigFile  saves the current object to a configuration file in Toml format
func (p *Project) SaveConfigFile() error {
	configPath := filepath.Join(p.ProjectPath, "config.toml")

	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(p)
	if err != nil {
		clean := p.Clean()
		if clean != nil {
			return clean
		}
		return err
	}
	err = ioutil.WriteFile(configPath, buf.Bytes(), 0660)
	if err != nil {
		clean := p.Clean()
		if clean != nil {
			return clean
		}
		return err
	}
	return nil
}

// Clean  makes sure all project  files in disc are safely removed
func (p *Project) Clean() error {
	if p.ProjectPath == "" {
		err := errors.New("Project path should not be empty")
		return err
	}
	err := os.RemoveAll(p.ProjectPath)
	if err != nil {
		return err
	}
	return nil
}

// Build run hugo on the root of project path to generate static files in public folder
// of the project path.
func (p *Project) Build() error {
	logThis.Info("Building.%s.", p.ProjectPath)
	sess := sh.NewSession().SetDir(p.ProjectPath)

	err := sess.Command("hugo").Run()
	if err != nil {
		beego.Info("Oopa, failed to build", err)
		clean := p.Clean()
		if clean != nil {
			return clean
		}
		return err
	}
	logThis.Info("*** done building ***")
	return err
}

// GenLorem populates pages with lorem ipsum for the page's content
func (p *Project) GenLorem() {
	logThis.Info("generating dummy data")
	for k, _ := range p.Pages {
		page := &p.Pages[k]
		buf := new(bytes.Buffer)
		head := new(bytes.Buffer)
		head.WriteString("## ")
		head.WriteString(lorem.Sentence(4, 7))
		head.WriteString("\n")
		buf.WriteString(head.String())
		n := []int{1, 2, 3, 4}
		for _, val := range n {
			buf.WriteString(lorem.Paragraph(val, val+1))
			buf.WriteString("\n\n")
		}

		page.Content = buf.String()
	}

	logThis.Info(" finished adding dummy content")
}

func (p *Project) Initialize(base string, name string, template string, theme string) error {
	err := initializeProject(p, base, name, template, theme)
	if err != nil {
		return err
	}
	return nil
}

func initializeProject(p *Project, base string, name string, template string, theme string) error {
	var projectsDir, templatesDir string
	projectsDir = beego.AppConfig.String("projectsDir")
	if projectsDir == "" {
		projectsDir = "projects"
	}

	templatesDir = beego.AppConfig.String("templatesDir")
	if templatesDir == "" {
		templatesDir = "templates"
	}
	b, _ := osext.ExecutableFolder()
	p.BaseDir = b
	if base != "" {
		p.BaseDir = base
	}

	p.Name = name
	p.Template = template
	if template == "" {
		p.Template = "default"
	}
	if theme == "" {
		p.Theme = "loraina"
	}
	sourceDir := filepath.Join(filepath.Join(p.BaseDir, templatesDir), p.Template)
	destDir := filepath.Join(filepath.Join(p.BaseDir, projectsDir), p.Name)

	src, err := os.Stat(sourceDir)
	if err != nil {
		return err
	}
	if !src.IsDir() {
		return errors.New("Oops, we cant get the template tou are asking")
	}
	_, err = os.Open(destDir)
	if !os.IsNotExist(err) {
		return errors.New("The project has already been built")
	}
	p.ProjectPath = destDir
	p.TemplatePath = sourceDir
	return nil
}

func (p *Project) InstallTemplate(name, theme string) error {
	err := installTemplate(p, name, theme)
	return err
}
func (p *Project) InstallTheme(theme string) error {
	err := copyTheme(p, theme)
	return err
}
func installTemplate(p *Project, templatename string, theme string) error {
	templatesDir := beego.AppConfig.String("templatesDir")
	if templatename == "" || p.Template == "" {
		templatename = "default"
	}
	if theme == "" || p.Theme == "" {
		theme = "loraina"
	}
	dest := p.ProjectPath
	sourceDir := filepath.Join(filepath.Join(p.BaseDir, templatesDir), templatename)
	err := cp.CopyDir(sourceDir, dest)
	if err != nil {
		if os.IsExist(err) {
			e := p.Clean()
			if e != nil {
				return e
			}
			e = cp.CopyDir(sourceDir, dest)
			if e != nil {
				return e
			}
		}
	}
	p.Template = templatename
	p.Theme = theme
	return nil
}
