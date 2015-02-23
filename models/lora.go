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

type Lora struct {
	Page       Page
	Project    Project
	Account    Account
	Section    Section
	SubSection SubSection
	Profile    Profile

	Pages       []Page
	Projects    []Project
	Accounts    []Account
	Sections    []Section
	SubSections []SubSection
	Profiles    []Profile

	HasPage       bool
	HasProject    bool
	HasAccount    bool
	HasSection    bool
	HasSubSection bool
	HasProfile    bool

	HasPages       bool
	HasProjects    bool
	HasAccounts    bool
	HasSections    bool
	HasSubSections bool
	HasProfiles    bool
	
	IsAdmin      bool
}

func (l *Lora) Add(v interface{}) {
	l.addObject(v)
}

func (l *Lora) addObject(v interface{}) {
	switch v.(type) {
	case Page:
		l.Page = v.(Page)
		l.HasPage = true
	case Project:
		l.Project = v.(Project)
		l.HasProject = true
	case Account:
		l.Account = v.(Account)
		l.HasAccount = true
		if l.Account.ClearanceLevel==6{
			l.IsAdmin=true
		}
	case Profile:
		l.Profile = v.(Profile)
		l.HasProfile = true
	case Section:
		l.Section = v.(Section)
		l.HasSection = true
	case SubSection:
		l.SubSection = v.(SubSection)
		l.HasSubSection = true

	case []Page:
		l.Pages = v.([]Page)
		l.HasPages = true
	case []Project:
		l.Projects = v.([]Project)
		l.HasProjects = true
	case []Account:
		l.Accounts = v.([]Account)
		l.HasAccounts = true
	case []Profile:
		l.Profiles = v.([]Profile)
		l.HasProfiles = true
	case []Section:
		l.Sections = v.([]Section)
		l.HasSections = true
	case []SubSection:
		l.SubSections = v.([]SubSection)
		l.HasSubSections = true
	default:
		logThis.Info("nothing matched")
	}

}
func NewLoraObject() *Lora {
	return &Lora{
		Pages:          []Page{},
		Projects:       []Project{},
		Accounts:       []Account{},
		HasAccount:     false,
		HasAccounts:    false,
		HasPage:        false,
		HasPages:       false,
		HasProject:     false,
		HasProjects:    false,
		HasProfile:     false,
		HasProfiles:    false,
		HasSection:     false,
		HasSections:    false,
		HasSubSections: false,
	}
}
