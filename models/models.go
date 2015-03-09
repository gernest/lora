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
	"html/template"

	"time"
)

type (
	// Account stores user information
	Account struct {
		Id              int64     `form:"-"`
		UserName        string    `toml:"-" sql:"not null;unique" form:"userName" valid:"Required"`
		Company         string    `toml:"company_name" sql:"unique" form:"company" valid:"Required"`
		Email           string    `toml:"email" sql:"unique" form:"email" valid:"Email"`
		Projects        []Project `toml:"-" form:"-"`
		Password        string    `toml:"-" form:"password" valid:"Required"`
		ConfirmPassword string    `toml:"-" sql:"-" form:"password2" valid:"Required"`
		Uploads         []Upload  `toml:"-" form:"-"`
		Profile         Profile   `toml:"-" form:"-"`
		ProfileId       int64     `toml:"-" form:"-"`
		ClearanceLevel  int       `toml:"-" form:"-"`
		CreatedAt       time.Time `toml:"-" form:"-"`
		UpdatedAt       time.Time `toml:"-" form:"-"`
		DeletedAt       time.Time `toml:"-" form:"-"`
	}

	// Project provide an abstract representation of a hugo project with other extra important details
	// about a website project
	Project struct {
		Id           int64
		Title        string    `toml:"title"`
		Name         string    `toml:"projectName" sql:"unique""`
		BaseDir      string    `toml:"-" sql:"-"`
		Template     string    `toml:"templateName"`
		Theme        string    `toml:"theme"`
		PublishDir   string    `toml:"publishDir"`
		ProjectPath  string    `toml:"-"`
		TemplatePath string    `toml:"-"`
		BaseUrl      string    `toml:"baseurl" sql:"-"`
		HomeUrl      string    `toml:"-" sql:"-"`
		LanguageCode string    `toml:"languageCode" sql:"-"`
		Pages        []Page    `toml:"pages"`
		AccountId    int64     `toml:"-" sql:"null"`
		Param        Param     `toml:"params"`
		ParamId      int64     `toml:"-"`
		Copyright    string    `toml:"copyright"`
		CreatedAt    time.Time `toml:"-"`
		UpdatedAt    time.Time `toml:"-"`
	}

	// Page representation of web page
	Page struct {
		Id          int64
		Title       string        `toml:"title"`
		ProjectId   int64         `toml:"-" `
		Content     string        `toml:"content" sql:"null;type:text"`
		ContentHtml template.HTML `toml:"-" sql:"-"`
		Slug        string        `toml:"slug" sql:"null"`
		Draft       bool          `toml:"draft" sql:"null"`
		Sections    []Section     `toml:"sections"`
		ContentPath string        `toml:"-" sql:"-"`
		CreatedAt   time.Time     `toml:"-"`
		UpdatedAt   time.Time     `toml:"-"`
	}

	//Section divides a page into small pieces
	Section struct {
		Id          int64
		PageId      int64
		Name        string        `toml:"name"`
		Title       string        `toml:"title"`
		Phone       string        `toml:"phone"`
		Photo       string        `toml:"photo"`
		Email       string        `toml:"email"`
		Address     string        `toml:"address"`
		Body        string        `toml:"body" sql:"null;type:text"`
		BodyHtml    template.HTML `sql:"-"`
		SubSections []SubSection  `toml:"subsections"`
		LastUpdate  string        `sql:"-" toml:"last_update"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	SubSection struct {
		Id         int64
		Name       string `toml:"name"`
		Photo      string `toml:"photo"`
		SectionId  int64
		Body       string        `toml:"body"`
		BodyHtml   template.HTML `sql:"-"`
		LastUpdate string        `sql:"-" toml:"last_update"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	Profile struct {
		Id        int64
		Phone     string
		Photo     string `sql:"null"`
		Thumbnail string
		UpdatedAt time.Time
		CreatedAT time.Time
	}

	Upload struct {
		Id   int64
		Path string
	}
	Param struct {
		Id              int64
		Slides          []Image `toml:"slides"`
		Author          string  `toml:"author"`
		Description     string  `toml:"description"`
		BackgroundImage string  `toml:"background"`
		Brand           string  `toml:"brand"`
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}

	Image struct {
		Id   int64
		Path string `toml:"path"`
	}
)
