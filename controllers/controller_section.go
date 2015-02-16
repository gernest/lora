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
package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/gernest/lora/models"
)

type SectonController struct {
	MainController
}

func (s *SectonController) Update() {
	sess := s.ActivateContent("section/edit")
	flash := beego.NewFlash()
	s.LayoutSections["JScripts"] = "jscript/editor.html"

	if sess == nil {
		flash.Error("you need to login inorder to update this page")
		flash.Store(&s.Controller)
		return
	}
	pageID, _ := s.GetInt64(":pageID")
	sectionID, _ := s.GetInt64(":sectionID")

	section := models.Section{}
	lora := models.NewLoraObject()
	subSections := []models.SubSection{}
	page := models.Page{}
	project := models.Project{}

	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		flash.Error("Whacko opening the database")
		flash.Store(&s.Controller)
		return
	}
	err = db.First(&section, sectionID).Error
	if err != nil {
		flash.Error("WHacko ", err)
		flash.Store(&s.Controller)
		return
	}
	if section.PageId != pageID {
		flash.Error("The page does not belong to this project")
		flash.Store(&s.Controller)
		return
	}
	db.Model(&section).Related(&subSections)
	for k := range subSections {
		n := &subSections[k]
		n.Sanitize()
	}
	if len(subSections) > 0 {
		lora.Add(subSections)
	}
	section.Sanitize()
	lora.Add(section)
	s.Data["lora"] = lora

	if s.Ctx.Input.Method() == "POST" {

		// TODO: cleanup this mess and add validation
		sectionContent := s.GetString("content")
		sectionTitle := s.GetString("sectionTitle")
		sectionName := s.GetString("sectionName")
		sectionPhone := s.GetString("sectionPhone")
		sectionEmail := s.GetString("sectionEmail")
		sectionAddress := s.GetString("sectionAddress")

		db.Find(&page, section.PageId)
		db.First(&project, page.ProjectId)

		uploadsDir := "projects/" + project.Name + "/static/img"

		if sectionTitle != "" {
			section.Title = sectionTitle
		}
		if sectionName != "" {
			section.Name = sectionName
		}

		if sectionPhone != "" {
			section.Phone = sectionPhone
		}
		if sectionEmail != "" {
			section.Email = sectionEmail
		}
		if sectionAddress != "" {
			section.Address = sectionAddress
		}

		// Handle File uploads
		_, fileHeader, err := s.GetFile("sectionPhoto")
		if err != nil {
			if err == http.ErrMissingFile {
				logThis.Info("There is no uploaded file")
			} else {
				logThis.Debug("Problem %v", err)
			}
		}
		if fileHeader != nil {
			logThis.Debug("Filename *%s* fileHead *%s*", fileHeader.Filename, fileHeader.Header)
			fileDestination := filepath.Join(uploadsDir, fileHeader.Filename)
			logThis.Debug("destination is %s", fileDestination)
			err = s.SaveToFile("sectionPhoto", fileDestination)

			if err != nil {
				logThis.Debug("Trouble Saving pic %v", err)
			}
			section.Photo = "/img/" + fileHeader.Filename
		}

		section.Body = sectionContent
		db.Save(&section)

		err = Rebuild(&page)
		if err != nil {
			flash.Error(" WHacko ", err)
			flash.Store(&s.Controller)
			return
		}
		s.Redirect("/accounts", 302)
	}
}
