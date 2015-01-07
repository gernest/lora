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
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/gernest/lora/models"
)

// AccountController hangles account related endpoints
type AccountController struct {
	MainController
}

// Index shows the accounts landing page
func (c *AccountController) Index() {
	sess := c.ActivateContent("accounts/home")
	c.LayoutSections["JScripts"] = "jscript/rest.html"
	c.SetNotice()

	flash := beego.NewFlash()
	lora:=models.NewLoraObject()
	if sess == nil {
		logThis.Info(" Attempt to access restircted page")
		flash.Error("You need to login to access this page")
		flash.Store(&c.Controller)
		c.Redirect("/accounts/login", 302)

	}
	a, err := checkUserByEmail(sess["email"].(string))
	if err != nil {
		logThis.Debug("User indentity trouble * %v", err)
		flash.Error("No such user")
		flash.Store(&c.Controller)
		return
	}
	lora.Add(a)
	c.Data["lora"]=lora
}

// Login athenticates the user
func (c *AccountController) Login() {

	sess := c.ActivateContent("accounts/login")
	c.SetNotice()
	if sess != nil {
		c.Redirect("/", 302)
	}
	if c.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		email := c.GetString("email")
		password := c.GetString("password")
		valid := validation.Validation{}
		valid.Email(email, "email")
		valid.Required(password, "password")
		if valid.HasErrors() {
			errormap := make(map[string]string)
			for _, err := range valid.Errors {
				errormap[err.Key] = err.Message
			}
			c.Data["Errors"] = errormap
			return
		}

		a, err := checkUserByEmail(email)
		if err != nil || a.Id == 0 {
			logThis.Debug(" %v", err)
			flash.Error("Sorry no  we have no record for this, try registering again or ask for help")
			flash.Store(&c.Controller)
			return
		}
		err = verifyPassword(&a, password)
		if err != nil {
			logThis.Debug("%v ", err)
			flash.Error("Incorrect Password")
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

		c.Redirect("/accounts", 302)
	}

}

// Logout deletes the session cookies
func (c *AccountController) Logout() {
	c.DelSession("xshabe")
	flash := beego.NewFlash()
	flash.Notice("You are now loged out")
	flash.Store(&c.Controller)
	c.Redirect("/", 302)
}

// Register creates a new account
func (c *AccountController) Register() {
	flash := beego.NewFlash()

	sess := c.ActivateContent("accounts/register")
	if sess != nil {
		logThis.Info("Session is still valid")
		flash.Notice("You have already registered an account")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
	}

	if c.Ctx.Input.Method() == "POST" {
		userName := c.GetString("userName")
		company := c.GetString("company")
		email := c.GetString("email")
		password := c.GetString("password")
		password2 := c.GetString("password2")
		terms := c.GetString("cb")
		if terms != "1" {
			flash.Error(" Please you need to accept terms of service in order to create a new account")
			flash.Store(&c.Controller)
			return
		}

		a := models.RegistrationForm{
			UserName: userName,
			Company:  company,
			Email:    email,
			Password: password,
			Confirm:  password2,
		}

		valid := validation.Validation{}
		if b, _ := valid.Valid(&a); !b {
			errormap := make(map[string]string)
			for _, err := range valid.Errors {
				s := strings.Split(err.Key, ".")
				errormap[s[0]] = err.Message
			}
			c.Data["Errors"] = errormap
			return
		}
		if password != password2 {
			flash.Error("Password Does not Match")
			flash.Store(&c.Controller)
			return
		}

		db, err := models.Conn()
		if err != nil {
			logThis.Info("%v", err)
			flash.Error("some fish opening database")
			flash.Store(&c.Controller)
			return
		}
		profile := models.Profile{
			Phone: "+27769000000",
		}
		account := models.Account{
			UserName: userName,
			Email:    email,
			Profile:  profile,
		}
		err = newAccountPassword(&account, password)
		if err != nil {
			logThis.Info("%v", err)
			flash.Error("some fish opening database")
			flash.Store(&c.Controller)
			return
		}

		query := db.Where("email= ?", account.Email).First(&account)
		if query.Error == nil {
			logThis.Debug("Trouble querying %s", query.Error.Error())
			flash.Error(email + "Already Registered")
			flash.Store(&c.Controller)
			return
		}
		db.Create(&account)
		if db.Error != nil {
			logThis.Info("%v", err)
			flash.Error("Sorry Internal Problems Occured, try again later")
			flash.Store(&c.Controller)
			return
		}
		flash.Notice("Your Account has been created successful you can login and enjoy")
		flash.Store(&c.Controller)
		c.Redirect("/accounts/login", 302)
	}

}

// Verify provides mechanism for account verification
func (c *AccountController) Verify() {}

// Profile shows and updates user infromation
func (c *AccountController) Profile() {}

// Remove deletes user account
func (c *AccountController) Remove() {}
