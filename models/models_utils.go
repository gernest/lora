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
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"

	"bitbucket.org/kardianos/osext"

	"github.com/gernest/lora/utils/logs"

	"github.com/astaxie/beego"
	cp "github.com/gernest/lora/utils/copy"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var logThis = logs.NewLoraLog()

func connStr() string {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_passwd")
	db_name := beego.AppConfig.String("db_name")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	dns = fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s", db_name, db_host, db_port, db_user, db_pass, db_sslmode)
	return dns
}

// Conn provied a *gorm.DB which is ready for interacting with database
func Conn() (*gorm.DB, error) {
	var dbPath, dialect string
	var err error
	var db gorm.DB
	dbPath = connStr()
	dialect = beego.AppConfig.String("db_dialect")
	db, err = gorm.Open(dialect, dbPath)
	//db.LogMode(true)
	return &db, err
}

func RunMigrations() {
	logThis.Info("Running Migrations")
	db, err := Conn()
	if err != nil {
		beego.Info("Failed to initialize database with error ", err)
		return
	}
	//db.LogMode(true)
	db.DropTableIfExists(new(Account))
	db.DropTableIfExists(new(Project))
	db.DropTableIfExists(new(Page))
	db.DropTableIfExists(new(Section))
	db.DropTableIfExists(new(SubSection))
	db.DropTableIfExists(new(Profile))
	db.AutoMigrate(new(Account), new(Profile), new(Project), new(Page), new(Section), new(SubSection))
	logThis.Success("migrations succeded")
}

func NewLoraProject(base string, name string, template string, theme string) (Project, error) {
	p := new(Project)
	err := p.Initialize(base, name, template, theme)
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}

	err = p.GenScaffold()
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}
	err = copyTheme(p, theme)
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}
	err = p.InitDir()
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}
	err = p.LoadConfigFile()
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}

	p.GenLorem() // fill content with dummy data

	return *p, nil
}

func GetAvailableThemes(base string) ([]string, error) {
	list, err := getResourceList("themes", base)
	return list, err
}
func GetAvailableTemplates(base string) ([]string, error) {
	list, err := getResourceList("templates", base)
	return list, err
}

func copyTheme(p *Project, name string) error {
	themeDir := beego.AppConfig.String("themesDir")
	if themeDir == "" {
		themeDir = "themes"
	}
	if name == "" {
		name = "loraina"
	}
	sourceDir := filepath.Join(filepath.Join(p.BaseDir, themeDir), name)
	destDir := filepath.Join(filepath.Join(p.ProjectPath, "themes"), name)
	err := cp.CopyDir(sourceDir, destDir)
	if err != nil {
		return err
	}
	p.Theme = name
	return nil
}

func sanitizeHTMLField(s string) []byte {
	unsafe := blackfriday.MarkdownCommon([]byte(s))
	safe := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return safe
}
func SanitizeTestHelper(s string) template.HTML {
	sane := sanitizeHTMLField(s)
	return template.HTML(sane)
}

func getFuncCall(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		_, filename := path.Split(file)
		s := fmt.Sprintf("[%s : %d] ", filename, line)
		return s
	}
	return ""
}

func getResourceList(s, location string) ([]string, error) {
	var (
		resource, resourcePath, resourceDir, base string
		resourceList                              []string
	)
	base = location
	if location == "" {
		baseD, err := osext.ExecutableFolder()
		if err != nil {
			return []string{}, err
		}
		base = baseD
	}

	switch s {
	case "themes":
		resource = "themes"
		resourceDir = beego.AppConfig.String("themesDir")
		if resourceDir == "" {
			resourceDir = resource
		}
	case "templates":
		resource = "templates"
		resourceDir = beego.AppConfig.String("templatesDIr")
		if resourceDir == "" {
			resourceDir = resource
		}
	}
	resourcePath = filepath.Join(base, resourceDir)
	if filepath.IsAbs(resourceDir) {
		resourcePath = resourceDir
	}
	file, err := os.Open(resourcePath)
	if err != nil {
		return []string{}, err
	}
	resourceList, err = file.Readdirnames(2)

	if err != nil {
		return []string{}, err
	}
	sort.Sort(sort.StringSlice(resourceList)) //Sort the list
	return resourceList, nil
}

func getLocalHost() string {
	port := beego.AppConfig.String("httpport")
	if port == "" {
		return ""
	}
	host := "localhost"
	scheme := fmt.Sprintf("http://%s:%s", host, port)
	return scheme
}
