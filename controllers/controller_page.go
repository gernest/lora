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
	"fmt"

	"github.com/astaxie/beego"
	"github.com/gernest/lora/models"
)

// PageController Ironically the user is not allowed to create any new page
type PageController struct {
	MainController
}

func (p *PageController) Preview() {
	projectID, err := p.GetInt64(":projectID")
	if err != nil {
		beego.Info("whacko ", err)
	}
	pageID, err := p.GetInt64(":pageID")
	if err != nil {
		beego.Info("whacko", err)
	}
	project := new(models.Project)
	page := new(models.Page)

	flash := beego.NewFlash()

	db, err := models.Conn()
	if err != nil {
		beego.Info("whacko", err)
	}
	err = db.First(project, projectID).Error
	if err != nil {
		beego.Info("whacko", err)
		flash.Error("Broken preview link")
		flash.Store(&p.Controller)
		p.Redirect("/accounts", 302)
	}
	err = db.First(page, pageID).Error
	if err != nil {
		beego.Info("whacko", err)
		flash.Error("Broken preview link")
		flash.Store(&p.Controller)
		p.Redirect("/accounts", 302)
	}
	link := "/preview/" + project.Name + "/" + page.Title
	p.Redirect(link, 302)
}

func (p *PageController) Update() {
	sess := p.ActivateContent("page/edit")
	flash := beego.NewFlash()
	lora := models.NewLoraObject()
	p.LayoutSections["JScripts"] = "jscript/editor.html"

	if sess == nil {
		flash.Error("you need to login inorder to update this page")
		flash.Store(&p.Controller)
		return
	}
	pageID, _ := p.GetInt64(":pageID")
	projectID, _ := p.GetInt64(":projectID")
	page := models.Page{}

	db, err := models.Conn()
	if err != nil {
		flash.Error("Whacko opening the database")
		flash.Store(&p.Controller)
		return
	}
	err = db.First(&page, pageID).Error
	if err != nil {
		flash.Error("WHacko ", err)
		flash.Store(&p.Controller)
		return
	}
	if page.ProjectId != projectID {
		flash.Error("The page does not belong to this project")
		flash.Store(&p.Controller)
		return
	}
	sections := []models.Section{}
	db.Model(&page).Related(&sections)
	for k := range sections {
		n := &sections[k]
		n.Sanitize()
	}
	lora.Add(sections)
	page.Sanitize()
	lora.Add(page)
	p.Data["lora"] = lora

	if p.Ctx.Input.Method() == "POST" {
		pageContent := p.GetString("content")
		page.Content = pageContent
		db.Save(&page)
		page.Sanitize()
		sp, _ := p.GetInt("saveAndPreview")
		err = Rebuild(&page)
		if err != nil {
			flash.Error("", err)
			flash.Store(&p.Controller)
			return
		}
		if sp == 1 {
			previewPath := fmt.Sprintf("/page/%d/%d/preview", page.ProjectId, page.Id)
			p.Redirect(previewPath, 302)
		}
		lora.Add(page)
		p.Data["lora"] = lora
	}
}
