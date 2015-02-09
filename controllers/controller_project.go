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
}

// NewProject Creates a new boilerplate hugo project and saves important data to database
func (p *ProjectController) NewProject() {

	var th, tmpl []Resource

	sess := p.ActivateContent("projects/new")

	themes, _ := models.GetAvailableThemes("")
	logThis.Debug("%v", themes)

	templates, _ := models.GetAvailableTemplates("")
	logThis.Debug("%v", templates)
	th = make([]Resource, len(themes))
	tmpl = make([]Resource, len(templates))
	for k, v := range themes {
		th[k] = Resource{Key: v, Value: v}
	}
	for k, v := range templates {
		tmpl[k] = Resource{Key: v, Value: v}
	}
	logThis.Info("Thems list %v", th)
	logThis.Info("templateList %v", tmpl)
	p.Data["themeList"] = &th
	p.Data["templateList"] = &tmpl

	if p.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		projectName := p.GetString("projectName")
		templeteName := p.GetString("templateName")
		themeName := p.GetString("themeName")
		logThis.Debug("Theme selected %s", themeName)
		logThis.Debug("template selected %s", templeteName)
		if sess == nil {
			flash.Error("You need  to login inorder to create a new site")
			flash.Store(&p.Controller)
			return
		}
		db, err := models.Conn()
		defer db.Close()
		if err != nil {
			logThis.Debug(":==> %v ", err)
			flash.Error("some fish opening database")
			flash.Store(&p.Controller)
			return
		}

		a, err := checkUserByEmail(sess["email"].(string))
		if err != nil {
			flash.Error("Sorry problem idenfying your acount please try again")
			flash.Store(&p.Controller)
			return
		}
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
		if db.NewRecord(project) {
			logThis.Debug("Failed to save into database %s", db.Error)
			flash.Error("Problem saving the project")
			flash.Store(&p.Controller)
			_ = project.Clean()
			return
		}
		beego.Info("Inital Build")
		err = project.Build()
		if err != nil {
			logThis.Info("Failed to Build %v", err)
			flash.Error("Failed to build project")
			flash.Store(&p.Controller)
			_ = project.Clean()
			return
		}

		// serve public folder as static
		staticPath := filepath.Join(project.ProjectPath, "www")
		previewPath := "/preview/" + project.Name
		deployPath := "/apps/" + project.Name

		beego.SetStaticPath(previewPath, staticPath) // preview
		beego.SetStaticPath(deployPath, staticPath)  // deploy local

		flash.Notice("your website has successful been created")
		flash.Store(&p.Controller)
		p.Redirect("/projects/list", 302)

	}

}

// Remove delets all project data from disc and database
func (p *ProjectController) Remove() {
	projectID, err := p.GetInt64(":id")
	if err != nil {
		beego.Info("some whacko %s", err)
	}
	beego.Info("project id is ", projectID)
	p.Data["projectId"] = projectID
	flash := beego.NewFlash()

	sess := p.ActivateContent("projects/delete")
	if p.Ctx.Input.Method() == "GET" {
		if sess == nil {
			flash.Error("You need  to login inorder to delete a site")
			flash.Store(&p.Controller)
			return
		}

	}
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

		em := sess["email"]
		a := models.Account{}
		a.Email = em.(string)
		query := db.Where("email= ?", a.Email).First(&a)
		if query.Error != nil {
			flash.Error("Sorry problem idenfying your acount please try again")
			flash.Store(&p.Controller)
			return
		}

		project := models.Project{}
		query = db.Model(&a).Related(&project)
		if project.Id != projectID || project.Name != projectName {
			flash.Error("project name mismatch  please try again with the correct name")
			flash.Store(&p.Controller)
			return
		}

		logThis.Info("deleting project %s", project.Name)

		// delete all pages
		pages := []models.Page{}
		query = db.Model(&project).Related(&pages)
		logThis.Event("deleting pages")
		for _, val := range pages {
			logThis.Event("deleting page *%s*", val.Title)
			db.Delete(&val)
		}
		logThis.Success("page deletion success")
		logThis.Event("deleting project from disc")
		err = project.Clean()
		if err != nil {
			flash.Error("Whaamy", err)
			flash.Store(&p.Controller)
			return
		}
		logThis.Event("Removing database records")
		err = db.Delete(&project).Error
		if err != nil {
			logThis.Debug(" WHammy %s", err)
			flash.Error("Whaamy")
			flash.Store(&p.Controller)
			return
		}
		logThis.Event("Updading user")
		db.Save(&a)
		logThis.Success("Project was deleted successful")
		flash.Notice("Your website has been deleted successful")
		flash.Store(&p.Controller)
		p.Redirect("/accounts", 302)

	}
}

// Preview redirects to the poject preview page, the pages are served as static files
func (p *ProjectController) Preview() {
	projectID, err := p.GetInt64(":id")
	if err != nil {
		beego.Info("Whaacko %s", err)
	}
	project := new(models.Project)

	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		beego.Info("Whacko whacko %s", err)
	}
	db.LogMode(true)
	db.First(project, projectID)

	previewLink := "/preview/" + project.Name
	p.Redirect(previewLink, 302)

}

// Update provides a restful project update
func (p *ProjectController) Update() {
	sess := p.ActivateContent("projects/update")

	flash := beego.NewFlash()
	lora := models.NewLoraObject()
	if sess == nil {
		flash.Error("You need  to login inorder to delete a site")
		flash.Store(&p.Controller)
		return
	}
	projectID, err := p.GetInt64(":id")
	if err != nil {
		beego.Info("Whaacko %s", err)
	}
	project := models.Project{}

	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		logThis.Debug("Whacko whacko %s", err)
		flash.Error("Sorry Internal problem")
		flash.Store(&p.Controller)
		return
	}
	db.LogMode(true)
	err = db.First(&project, projectID).Error
	if err != nil {
		logThis.Debug("Whacko whacko %s", err)
		flash.Error("Sorry Internal problem")
		flash.Store(&p.Controller)
		return
	}
	pages := []models.Page{}
	err = db.Model(&project).Related(&pages).Error
	if err != nil {
		logThis.Debug("Whacko whacko %s", err)
		flash.Error("Sorry Internal problem")
		flash.Store(&p.Controller)
		return
	}
	param:=models.Param{}
	db.First(&param, project.ParamId)
	
	project.Param=param
	lora.Add(pages)
	lora.Add(project)
	p.Data["lora"] = lora
	
	if p.Ctx.Input.Method()=="POST" {
		projectTitle:=p.GetString("projectTitle")
		paramsAuthor:=p.GetString("paramAuthor")
		paramDescription:=p.GetString("paramDescription")
		paramBrand:=p.GetString("paramBrand")
		
		err=project.LoadConfigFile()
		if err!=nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}
				
		if projectTitle!="" {
			project.Title=projectTitle
		}
		if paramDescription!="" {
			project.Param.Description=paramDescription
		}
		if paramsAuthor!="" {
			project.Param.Author=paramsAuthor
		}
		if paramBrand!="" {
			project.Param.Brand=paramBrand
		}
		
		err=db.Save(&project).Error
		if err!=nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}
			
		err=project.SaveConfigFile()
		if err!=nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}
			
		err=project.Build()
		if err!=nil {
			flash.Error("Fuck %s", err)
			flash.Store(&p.Controller)
			return
		}
		
		p.Redirect("/accounts", 302)		
	}
}

// Deploy prepares and pushes the project to the cloud
// TODO
func (p *ProjectController) Deploy() {
	p.ActivateView("notyet")
}

func (p *ProjectController) List() {
	sess := p.ActivateContent("projects/list")
	p.SetNotice()
	flash := beego.NewFlash()
	lora := models.NewLoraObject()
	if sess == nil {
		flash.Error("You need to login to access this page")
		flash.Store(&p.Controller)
		beego.Info("Session not set yet")
		p.Redirect("/accounts/login", 302)

	}
	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		beego.Info(":==> ", err)
		flash.Error("If you see this message, please report it by sending us aa email")
		flash.Store(&p.Controller)
		return
	}

	a := models.Account{}
	a.Email = sess["email"].(string)
	query := db.Where("email= ?", a.Email).First(&a)
	if query.Error != nil {
		flash.Error("Sorry problem idenfying your acount please try again")
		flash.Store(&p.Controller)
		return
	}

	projects := []models.Project{}
	db.Model(&a).Related(&projects)
	lora.Add(projects)

	if p.Ctx.Input.IsAjax() {
		p.Data["json"] = lora
		logThis.Info("AJAX Request")
		p.ServeJson()
	}
	p.Data["lora"] = lora
}

func(p *ProjectController)Download(){
	
}