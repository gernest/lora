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
	"github.com/astaxie/beego"
	"github.com/gernest/lora/models"
)

type SectonController struct {
	MainController
}

func (s *SectonController) Update() {
	sess := s.ActivateContent("section/edit")
	flash := beego.NewFlash()
	lora := models.NewLoraObject()
	s.LayoutSections["JScripts"] = "jscript/editor.html"

	if sess == nil {
		flash.Error("you need to login inorder to update this page")
		flash.Store(&s.Controller)
		return
	}
	pageID, _ := s.GetInt64(":pageID")
	sectionID, _ := s.GetInt64(":sectionID")
	section := models.Section{}

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
	subSections := []models.SubSection{}
	db.Model(&section).Related(subSections)
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
		sectionContent := s.GetString("content")
		section.Body = sectionContent
		db.Save(&section)
		page := new(models.Page)
		db.Find(page, section.PageId)
		err = Rebuild(page)
		if err != nil {
			flash.Error(" WHacko ", err)
			flash.Store(&s.Controller)
			return
		}
		s.Redirect("/accounts", 302)
	}
}
