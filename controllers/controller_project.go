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
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/gernest/lora/models"
)

// ProjectController for lora projects
type ProjectController struct {
	MainController
}
type Resource struct {
	Key   string
	Value string
	Image string
}

// NewProject Creates a new boilerplate hugo project and saves important data to database
func (p *ProjectController) NewProject() {

	sess := p.ActivateContent("projects/new")

	// Get available themes and templates
	themes, _ := models.GetAvailableThemes("")
	templates, _ := models.GetAvailableTemplates("")

	p.Data["themeList"] = &themes
	p.Data["templateList"] = &templates

	if p.Ctx.Input.Method() == "POST" {

		flash := beego.NewFlash()
		projectName := p.GetString("projectName")
		templeteName := p.GetString("templateName")
		themeName := p.GetString("themeName")

		if sess == nil {
			flash.Error("You need  to login inorder to create a new site")
			flash.Store(&p.Controller)
			return
		}

		db, err := models.Conn()
		defer db.Close()
		if err != nil {
			logThis.Debug(" %v ", err)
			flash.Error("There is a Problem, Try again")
			flash.Store(&p.Controller)
			return
		}

		a := sess["account"].(*models.Account)

		// Create a new project
		project, err := models.NewLoraProject("", projectName, templeteName, themeName)
		if err != nil {
			logThis.Critical("Failed **%v**", err)
			flash.Error("failed to create the project")
			flash.Store(&p.Controller)
			return
		}

		logThis.Info("saving project to database")
		project.AccountId = a.Id
		db.Create(&project)

		if db.Error != nil {
			logThis.Info("holly shit check this mess %v", db.Error)
			flash.Error("some fish happened")
			flash.Store(&p.Controller)
			return
		}
		ps := &project
		err = ps.GenContent()
		if err != nil {
			logThis.Debug("holly shit check this mess %s", err.Error())
			flash.Error("some fish happened")
			flash.Store(&p.Controller)
			return
		}
		project.SetBaseUrl()
		err = ps.SaveConfigFile()
		if err != nil {
			logThis.Debug("holly shit check this mess %s", db.Error)
			flash.Error("some fish happened sorry")
			flash.Store(&p.Controller)
			return
		}

		db.Save(project)

		// Build the project
		logThis.Info("Inital Build")
		err = project.Build()
		if err != nil {
			logThis.Info("Failed to Build %v", err)
			flash.Error("Failed to build project")
			flash.Store(&p.Controller)
			_ = project.Clean()
			return
		}

		// The following is an attempt to make the project preview possible
		// I use the dymanic setting of static paths, to expose the www folder
		// The project build site is inside the www folder
		staticPath := filepath.Join(project.ProjectPath, "www")
		previewPath := "/preview/" + project.Name
		deployPath := "/apps/" + project.Name

		beego.SetStaticPath(previewPath, staticPath) // preview
		beego.SetStaticPath(deployPath, staticPath)  // deploy local

		flash.Notice("your website has successful been created")
		flash.Store(&p.Controller)
		p.Redirect("/web/project/list", 302)
	}

}

func (p *ProjectController) Remove() {
	flash := beego.NewFlash()

	sess := p.ActivateContent("projects/delete")
	if p.Ctx.Input.Method() == "GET" {
		if sess == nil {
			flash.Error("You need  to login inorder to delete a site")
			flash.Store(&p.Controller)
			return
		}

	}

	projectID, err := p.GetInt64(":id")
	if err != nil {
		logThis.Info("some whacko %s", err)
	}
	logThis.Info("project id is ", projectID)
	p.Data["projectId"] = projectID

	if p.Ctx.Input.Method() == "POST" {

		projectName := p.GetString("projectName")
		if sess == nil {
			flash.Error("You need  to login inorder to delete a site")
			flash.Store(&p.Controller)
			return
		}
		db, err := models.Conn()
		defer db.Close()
		if err != nil {
			beego.Info(":==> ", err)
			flash.Error("some fish opening database")
			flash.Store(&p.Controller)
			return
		}

		a := sess["account"].(*models.Account)

		project := models.Project{}
		db.Model(a).Related(&project)
		if project.Id != projectID || project.Name != projectName {
			flash.Error("project name mismatch  please try again with the correct name")
			flash.Store(&p.Controller)
			return
		}

		logThis.Info("deleting project %s", project.Name)

		logThis.Event("deleting project from disc")
		err = project.Clean()
		if err != nil {
			flash.Error("Whaamy", err)
			flash.Store(&p.Controller)
			return
		}

		logThis.Event("Removing project from database")
		// delete all pages, sections and subsections
		pages := []models.Page{}
		paramsProject := []models.Param{}

		db.Model(&project).Related(&pages)

		logThis.Event("deleting pages")
		for _, val := range pages {
			sections := []models.Section{}
			db.Model(&val).Related(&sections)
			logThis.Event("deleting page *%s*", val.Title)
			db.Delete(&val)

			for _, section := range sections {
				subsections := []models.SubSection{}
				db.Model(&section).Related(&subsections)
				db.Delete(&section)

				for _, sub := range subsections {
					db.Delete(sub)
				}
			}
		}

		// Deleting params
		db.First(&paramsProject, project.ParamId)

		db.Delete(&paramsProject)

		err = db.Delete(&project).Error
		if err != nil {
			logThis.Debug(" WHammy %s", err)
			flash.Error("Whaamy")
			flash.Store(&p.Controller)
			return
		}

		// Update user
		logThis.Event("Updading user")
		db.Save(a)

		logThis.Success("Project was deleted successful")
		flash.Notice("Your website has been deleted successful")
		flash.Store(&p.Controller)
		p.Redirect("/web/accounts", 302)

	}
}

// Preview redirects to the poject preview page, the pages are served as static files
func (p *ProjectController) Preview() {
	projectID, err := p.GetInt64(":id")
	if err != nil {
		logThis.Info("Whaacko %s", err)
	}

	project := models.Project{}

	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		beego.Info("Whacko whacko %s", err)
	}
	db.First(&project, projectID)

	previewLink := "/preview/" + project.Name
	p.Redirect(previewLink, 302)

}

// Update provides a restful project update
func (p *ProjectController) Update() {
	flash := beego.NewFlash()
	sess := p.ActivateContent("projects/update")
	if sess == nil {
		flash.Error("You need  to login inorder to delete a site")
		flash.Store(&p.Controller)
		return
	}

	projectID, err := p.GetInt64(":id")
	if err != nil {
		beego.Info("Whaacko %s", err)
	}

	lora := models.NewLoraObject()
	project := models.Project{}
	pages := []models.Page{}
	param := models.Param{}

	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		logThis.Debug("Whacko whacko %s", err)
		flash.Error("Sorry Internal problem")
		flash.Store(&p.Controller)
		return
	}

	err = db.First(&project, projectID).Error
	if err != nil {
		logThis.Debug("Whacko whacko %s", err)
		flash.Error("Sorry Internal problem")
		flash.Store(&p.Controller)
		return
	}

	err = db.Model(&project).Related(&pages).Error
	if err != nil {
		logThis.Debug("Whacko whacko %s", err)
		flash.Error("Sorry Internal problem")
		flash.Store(&p.Controller)
		return
	}

	db.First(&param, project.ParamId)

	project.Param = param
	lora.Add(pages)
	lora.Add(project)
	p.Data["lora"] = lora

	if p.Ctx.Input.Method() == "POST" {

		// TODO: cleanup this mess and add validtion
		projectTitle := p.GetString("projectTitle")
		paramsAuthor := p.GetString("paramAuthor")
		paramDescription := p.GetString("paramDescription")
		paramBrand := p.GetString("paramBrand")

		err = project.LoadConfigFile()
		if err != nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}

		if projectTitle != "" {
			project.Title = projectTitle
		}
		if paramDescription != "" {
			project.Param.Description = paramDescription
		}
		if paramsAuthor != "" {
			project.Param.Author = paramsAuthor
		}
		if paramBrand != "" {
			project.Param.Brand = paramBrand
		}

		err = db.Save(&project).Error
		if err != nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}

		err = project.SaveConfigFile()
		if err != nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}

		err = project.Build()
		if err != nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}

		p.Redirect("/web/accounts", 302)
	}
}

// Deploy prepares and pushes the project to the cloud
// TODO
func (p *ProjectController) Deploy() {
	p.ActivateView("notyet")
}

func (p *ProjectController) List() {
	flash := beego.NewFlash()
	sess := p.ActivateContent("projects/list")
	p.SetNotice()
	if sess == nil {
		flash.Error("You need to login to access this page")
		flash.Store(&p.Controller)
		beego.Info("Session not set yet")
		p.Redirect("/accounts/login", 302)
		return

	}
	a := sess["account"].(*models.Account)

	lora := models.NewLoraObject()
	projects := []models.Project{}

	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		beego.Info(":==> ", err)
		flash.Error("If you see this message, please report it by sending us aa email")
		flash.Store(&p.Controller)
		return
	}
	db.Model(a).Related(&projects)
	lora.Add(projects)
	p.Data["lora"] = lora
}
