package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/validation"

	"github.com/astaxie/beego"
	"github.com/gernest/lora/models"
)

type ProfileController struct {
	MainController
}

func (p *ProfileController) Edit() {
	sess := p.ActivateContent("profile/edit")
	flash := beego.NewFlash()
	p.LayoutSections["JScripts"] = "jscript/editor.html"

	if sess == nil {
		flash.Error("you need to login inorder to update this page")
		flash.Store(&p.Controller)
		return
	}

	profileID, _ := p.GetInt64(":profileID")
	profile := models.Profile{}

	db, err := models.Conn()
	if err != nil {
		flash.Error("Whacko opening the database")
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

	err = db.First(&profile, profileID).Error
	if err != nil {
		flash.Error("WHacko ", err)
		flash.Store(&p.Controller)
		return
	}
	if a.ProfileId != profileID {
		flash.Error("ou are no authorized to view this page")
		flash.Store(&p.Controller)
		return
	}
	p.Data["user"] = &a
	p.Data["profile"] = &profile

	if p.Ctx.Input.Method() == "POST" {
		company := p.GetString("company")
		phone := p.GetString("phone")
		uForm := models.UserForm{Company: company}
		uProfile := models.UserProfileForm{Phone: phone}

		v1 := validation.Validation{}
		v2 := validation.Validation{}
		errMap := make(map[string]string)
		if b, _ := v1.Valid(&uForm); !b {
			for _, err := range v1.Errors {
				s := strings.Split(err.Key, ".")
				fmt.Println(s)
				errMap[s[0]] = err.Message
			}
		}
		if b, _ := v2.Valid(&uProfile); !b {
			for _, err := range v2.Errors {
				s := strings.Split(err.Key, ".")
				errMap[s[0]] = err.Message
			}
		}
		if len(errMap) != 0 {
			p.Data["Errors"] = errMap
			return
		}
		a.Company = company
		profile.Phone = phone
		db.Save(&a)
		db.Save(&profile)

		rd := fmt.Sprintf("/profile/%d/view", profile.Id)
		p.Redirect(rd, 302)
	}
}

func (p *ProfileController) Display() {
	sess := p.ActivateContent("profile/display")
	flash := beego.NewFlash()
	p.LayoutSections["JScripts"] = "jscript/editor.html"

	if sess == nil {
		flash.Error("you need to login inorder to update this page")
		flash.Store(&p.Controller)
		return
	}

	profileID, _ := p.GetInt64(":profileID")
	profile := models.Profile{}

	db, err := models.Conn()
	if err != nil {
		flash.Error("Whacko opening the database")
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

	err = db.First(&profile, profileID).Error
	if err != nil {
		flash.Error("WHacko ", err)
		flash.Store(&p.Controller)
		return
	}
	if a.ProfileId != profileID {
		flash.Error("ou are no authorized to view this page")
		flash.Store(&p.Controller)
		return
	}
	p.Data["user"] = &a
	p.Data["profile"] = &profile
}
