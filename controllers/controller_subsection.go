package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gernest/lora/models"
)

type SubSectionController struct {
	MainController
}

func (s SubSectionController) Update() {
	sess := s.ActivateContent("section/edit")
	flash := beego.NewFlash()
	s.LayoutSections["JScripts"] = "jscript/editor.html"

	if sess == nil {
		flash.Error("you need to login inorder to update this page")
		flash.Store(&s.Controller)
		return
	}
	subSectionID, _ := s.GetInt64(":subsectionID")
	sectionID, _ := s.GetInt64(":sectionID")
	subSection := models.SubSection{}

	db, err := models.Conn()
	if err != nil {
		flash.Error("Whacko opening the database")
		flash.Store(&s.Controller)
		return
	}
	err = db.First(&subSection, subSectionID).Error
	if err != nil {
		flash.Error("WHacko ", err)
		flash.Store(&s.Controller)
		return
	}
	if subSection.SectionId != sectionID {
		flash.Error("The page does not belong to this project")
		flash.Store(&s.Controller)
		return
	}
	subSection.Sanitize()
	s.Data["subSection"] = subSection

	if s.Ctx.Input.Method() == "POST" {
		subSectionContent := s.GetString("content")
		subSection.Body = subSectionContent
		db.Save(&subSection)

		s.Redirect("/accounts", 302)
	}
}
