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
	"github.com/gernest/lora/utils/logs"
)

var logThis = logs.NewLoraLog()

// MainController provides base methods for all lora controllers
type MainController struct {
	beego.Controller
}

// ActivateContent makes it easy to add layout to templates, it also checks
// Session cookie if is set and do the initializing stuffs
func (c *MainController) ActivateContent(view string) map[string]interface{} {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHeader"] = "header.html"
	c.LayoutSections["HtmlFooter"] = "footer.html"
	c.TplNames = view + ".html"
	c.Layout = "base.html"

	logThis.Info("Checking session")
	sess := c.GetSession("xshabe")
	if sess != nil {
		m := sess.(map[string]interface{})
		c.Data["Username"] = m["username"]
		logThis.Success("Session found *%v*", m["username"])

		db, err := models.Conn()
		defer db.Close()
		em := m["email"]
		a := models.Account{}
		a.Email = em.(string)
		err = db.Where("email= ?", a.Email).First(&a).Error
		if err != nil {
			return nil
		}
		prof := models.Profile{}
		err = db.First(&prof, a.ProfileId).Error
		if err != nil {
			return nil
		}

		c.Data["InSession"] = 1
		c.Data["avatar"] = &prof
		c.Data["acc"] = &a
		return m

	}
	logThis.Warning("No session found")
	return nil
}

// Get takes you home baby
func (c *MainController) Get() {
	_ = c.ActivateContent("index")
	c.SetNotice()
}

// Notice this is an old school notice page
func (c *MainController) Notice() {
	_ = c.ActivateContent("notice")

	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		c.Data["notice"] = n
	}

}

// SetNotice makes it easier to set flash notices
func (c *MainController) SetNotice() {
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		c.Data["notice"] = n
	}
}

func (c *MainController) ActivateView(view string) {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHeader"] = "header.html"
	c.LayoutSections["HtmlFooter"] = "footer.html"
	c.TplNames = view + ".html"
	c.Layout = "base.html"
}
