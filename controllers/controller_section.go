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
	s.Data["subSections"] = subSections
	section.Sanitize()
	s.Data["section"] = section

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
