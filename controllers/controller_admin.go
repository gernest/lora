package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gernest/lora/models"
)

type AdminController struct {
	MainController
}

func (c *AdminController) Index() {
	sess := c.ActivateContent("admin/home")
	flash := beego.NewFlash()
	if sess == nil {
		flash.Error("Need to login")
		flash.Store(&c.Controller)
		return
	}
	accounts := []models.Account{}
	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		logThis.Info("%v", err)
		flash.Error("some fish opening database")
		flash.Store(&c.Controller)
		return
	}
	db.Find(&accounts)
	for key := range accounts {
		v := &accounts[key]
		projects := []models.Project{}
		profile := models.Profile{}

		// Check profile
		db.First(&profile, v.ProfileId)

		// Check projects
		db.Model(v).Related(&projects)

		// Attach objects to the account
		v.Profile = profile
		v.Projects = projects
	}
	c.Data["users"] = &accounts
}

func (c *AdminController) CreateAccounts() {}

func (c *AdminController) DeleteAccounts() {}

func (c *AdminController) UpdateAccounts() {}
