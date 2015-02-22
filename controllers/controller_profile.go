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
	"net/http"
	"path/filepath"
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
	defer db.Close()
	if err != nil {
		flash.Error("Whacko opening the database")
		flash.Store(&p.Controller)
		return
	}
	a := sess["account"].(*models.Account)

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
		uploadsDir := beego.AppConfig.String("uploadsDir")

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
		
		// Handle profile picture upload
		// If no file is chosen log and ignore
		// returning other errors
		_, fileHeader, err := p.GetFile("profilePicture")
		if err != nil {
			if err == http.ErrMissingFile {
				logThis.Info("There is no uploaded file")
			} else {
				errMap["profilePic"] = err.Error()

			}
		}
		if fileHeader != nil {
			logThis.Debug("Filename *%s* fileHead *%s*", fileHeader.Filename, fileHeader.Header)
			fileDestination := filepath.Join(uploadsDir, fileHeader.Filename)
			logThis.Debug("destination is %s", fileDestination)
			err = p.SaveToFile("profilePicture", fileDestination)

			if err != nil {
				logThis.Debug("Trouble Saving pic %v", err)
			}
			profile.Photo = "/" + fileDestination
		}

		if len(errMap) != 0 {
			p.Data["Errors"] = errMap
			return
		}
		a.Company = company
		profile.Phone = phone
		db.Save(a)
		db.Save(&profile)
		
		// Build a url leading back to profile view page
		profileViewPath := fmt.Sprintf("/web/accounts/profile/%d/view", profile.Id)
		p.Redirect(profileViewPath, 302)
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
	defer db.Close()
	if err != nil {
		flash.Error("Whacko opening the database")
		flash.Store(&p.Controller)
		return
	}
	a :=sess["account"].(*models.Account)
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
	lora := models.NewLoraObject()
	lora.Add(*a)
	lora.Add(profile)
	p.Data["user"] = a
	p.Data["profile"] = &profile
	p.Data["lora"] = lora
}
