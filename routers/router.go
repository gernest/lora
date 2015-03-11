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

package routers

import (
	"github.com/astaxie/beego"
	"github.com/gernest/lora/controllers"
	"github.com/gernest/lora/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/notice", &controllers.MainController{}, "get:Notice")

	ns := beego.NewNamespace("/web",
		beego.NSBefore(filters.AuthLevelSix),
		beego.NSNamespace("/accounts",
			beego.NSRouter("/", &controllers.AccountController{}, "*:Index"),
			beego.NSRouter("/register", &controllers.AccountController{}, "get,post:Register"),
			beego.NSRouter("/login", &controllers.AccountController{}, "get,post:Login"),
			beego.NSRouter("/logout", &controllers.AccountController{}, "get:Logout"),
			beego.NSNamespace("/profile",
				beego.NSRouter("/:profileID:int/edit", &controllers.ProfileController{}, "get,post:Edit"),
				beego.NSRouter("/:profileID:int/view", &controllers.ProfileController{}, "get:Display"),
			),
			beego.NSNamespace("/files",
				beego.NSRouter("/upload", &controllers.FilesController{}, "get,post:Upload"),
			),
		),
		beego.NSNamespace("/admin",
			beego.NSRouter("/", &controllers.AdminController{}, "get:Index"),
		),
		beego.NSNamespace("/project",
			beego.NSRouter("/new", &controllers.ProjectController{}, "get,post:NewProject"),
			beego.NSRouter("/:id:int/update/", &controllers.ProjectController{}, "get,post:Update"),
			beego.NSRouter("/:id:int/delete", &controllers.ProjectController{}, "get,post:Remove"),
			beego.NSRouter("/:id:int/preview", &controllers.ProjectController{}, "get,post:Preview"),
			beego.NSRouter("/:id:int/deploy", &controllers.ProjectController{}, "get,post:Deploy"),
			beego.NSRouter("/list", &controllers.ProjectController{}, "get:List"),
		),
		beego.NSNamespace("/page",
			beego.NSRouter("/:projectID:int/:pageID:int/update", &controllers.PageController{}, "get,post:Update"),
			beego.NSRouter("/:projectID:int/:pageID:int/preview", &controllers.PageController{}, "get,post:Preview"),
		),
		beego.NSNamespace("/section",
			beego.NSRouter("/:pageID:int/:sectionID:int/update", &controllers.SectonController{}, "get,post:Update"),
		),
		beego.NSNamespace("/subsection",
			beego.NSRouter("/subsection/:sectionID:int/:subSectionID:int/update", &controllers.SubSectionController{}, "get,post:Update"),
		),
	)

	nsLora := beego.NewNamespace("/site",
		beego.NSNamespace("/pages",
			beego.NSRouter("/pricing", &controllers.LoraController{}, "get:Pricing"),
			beego.NSRouter("/services", &controllers.LoraController{}, "get:Services"),
			beego.NSRouter("/contacts", &controllers.LoraController{}, "get:Contacts"),
			beego.NSRouter("/terms", &controllers.LoraController{}, "get:Terms"),
			beego.NSRouter("/legal", &controllers.LoraController{}, "get:Legal"),
			beego.NSRouter("/help", &controllers.LoraController{}, "get:Help"),
			beego.NSRouter("/companies", &controllers.LoraController{}, "get:Companies"),
		),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(nsLora)

	beego.InsertFilter("/", beego.BeforeRouter, filters.StaticProxy)
	beego.InsertFilter("/*", beego.BeforeRouter, filters.StaticProxy)
}
