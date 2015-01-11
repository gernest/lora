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
		Id             int64     `toml:"-"`
		UserName       string    `toml:"-" sql:"not null;unique"`
		Company        string    `toml:"company_name" sql:"unique"`
		Email          string    `toml:"email" sql:"unique"`
		Projects       []Project `toml:"-"`
		Password       string    `toml:"-"`
		Uploads        []Upload  `toml:"-"`
		Profile        Profile   `toml:"-"`
		ProfileId      int64     `toml:"-"`
		ClearanceLevel int       `toml:"-"`
		CreatedAt      time.Time `toml:"-"`
		UpdatedAt      time.Time `toml:"-"`
		DeletedAt      time.Time `toml:"-"`
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
		CreatedAt    time.Time `toml:"-"`
		UpdatedAt    time.Time `toml:"-"`
	}

	// Page representation of web page
	Page struct {
		Id          int64
		Title       string        `toml:"title"`
		ProjectId   int64         `toml:"-" `
		Content     string        `toml:"content" sql:"null;type:text"`
		ContentHtml template.HTML `toml:"-" sql:"-""`
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
		Body        string        `toml:"body" sql:"null;type:text"`
		BodyHtml    template.HTML `sql:"-"`
		Pre         string
		Pro         string
		SubSections []SubSection
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	SubSection struct {
		Id        int64
		SectionId int64
		Body      string
		BodyHtml  template.HTML `sq;"-"`
		Pre       string
		Pro       string
	}
	Profile struct {
		Id        int64
		Phone     string
		Photo     string `sql:"null`
		Thumbnail string
		UpdatedAt time.Time
		CreatedAT time.Time
	}

	Upload struct {
		Id   int64
		Path string
	}
)
