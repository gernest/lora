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
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/gernest/lora/models"
)

type AccountController struct {
	MainController
}

func (c *AccountController) Index() {
	sess := c.ActivateContent("accounts/home")
	c.SetNotice()

	flash := beego.NewFlash()
	if sess == nil {
		flash.Error("You tried to access restricted page without permission")
		flash.Store(&c.Controller)
		c.Redirect("/accounts/login", 302)
		return
	}
	a := sess["account"].(*models.Account)
	c.Data["user"] = a
	c.Data["Title"] = "Account"
}

func (c *AccountController) Login() {

	sess := c.ActivateContent("accounts/login")
	c.SetNotice()
	if sess != nil {
		c.Redirect("/", 302)
	}
	c.Data["Title"] = "login"
	if c.Ctx.Input.Method() == "POST" {
		loginForm := models.LoginForm{}
		flash := beego.NewFlash()

		if err := c.ParseForm(&loginForm); err != nil {
			logThis.Debug("%v", err)
		}
		valid := validation.Validation{}

		if b, _ := valid.Valid(&loginForm); !b {
			errMap := make(map[string]string)
			for _, v := range valid.Errors {
				errMap[v.Field] = v.Message
				logThis.Dump(v)
			}
			c.Data["FormErrors"] = &errMap
			return
		}
		a, err := checkUserByEmail(loginForm.Email)
		if err != nil || a.Id == 0 {
			logThis.Debug(" %v", err)
			flash.Error("Sorry wrong username or password")
			flash.Store(&c.Controller)
			return
		}
		err = verifyPassword(&a, loginForm.Password)
		if err != nil {
			logThis.Debug("%v ", err)
			flash.Error("Wrong username or password")
			flash.Store(&c.Controller)
			return
		}

		// Creating a session and going to the home page
		m := make(map[string]interface{})
		m["username"] = a.UserName
		m["email"] = a.Email
		m["timestamp"] = time.Now()
		notice := fmt.Sprintf("Welcome   %s", a.UserName)
		c.SetSession("xshabe", m)

		flash.Notice(notice)
		flash.Store(&c.Controller)

		c.Redirect("/web/accounts", 302)
	}

}

func (c *AccountController) Logout() {
	c.DelSession("xshabe")
	flash := beego.NewFlash()
	flash.Notice("You are now loged out")
	flash.Store(&c.Controller)
	c.Redirect("/", 302)
}

func (c *AccountController) Register() {
	flash := beego.NewFlash()

	sess := c.ActivateContent("accounts/register")
	if sess != nil {
		flash.Notice("You have already registered an account")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	c.Data["Title"] = "signup"

	if c.Ctx.Input.Method() == "POST" {
		usr := models.Account{}
		if err := c.ParseForm(&usr); err != nil {
			logThis.Debug("%v", err)
		}
		terms := c.GetString("cb")
		if terms != "1" {
			flash.Error(" Please you need to accept terms of service in order to create a new account")
			flash.Store(&c.Controller)
			return
		}
		valid := validation.Validation{}
		if b, _ := valid.Valid(&usr); !b {
			errMap := make(map[string]string)
			for _, v := range valid.Errors {
				errMap[v.Field] = v.Message
				logThis.Dump(v)
			}
			c.Data["FormErrors"] = &errMap
			return
		}
		if usr.Password != usr.ConfirmPassword {
			flash.Error("Password Does not Match")
			flash.Store(&c.Controller)
			return
		}

		db, err := models.Conn()
		defer db.Close()
		if err != nil {
			logThis.Info("%v", err)
			flash.Error("some fish opening database")
			flash.Store(&c.Controller)
			return
		}
		profile := models.Profile{
			Phone: "+27769000000",
		}
		usr.ClearanceLevel = 6
		usr.Profile = profile

		if err = newAccountPassword(&usr, usr.Password); err != nil {
			logThis.Info("%v", err)
			flash.Error("some fish opening database")
			flash.Store(&c.Controller)
			return
		}
		pf := &usr.Profile
		if err = pf.GenerateIdenticon("", usr.UserName); err != nil {
			logThis.Debug("%v", err)
		}

		if err = db.Create(&usr).Error; err != nil {
			logThis.Info("%v", err)
			flash.Error("Sorry Internal Problems Occured, try again later")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Your Account has been created successful you can login and enjoy")
		flash.Store(&c.Controller)
		c.Redirect("/web/accounts/login", 302)
	}

}
