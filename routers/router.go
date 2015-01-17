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

	// Accounts
	beego.Router("/accounts", &controllers.AccountController{}, "*:Index")
	beego.Router("/accounts/register", &controllers.AccountController{}, "get,post:Register")
	beego.Router("/accounts/login", &controllers.AccountController{}, "get,post:Login")
	beego.Router("/accounts/logout", &controllers.AccountController{}, "get:Logout")

	// Profile
	beego.Router("/profile/:profileID:int/edit", &controllers.ProfileController{}, "get,post:Edit")
	beego.Router("/profile/:profileID:int/view", &controllers.ProfileController{}, "get:Display")

	// Projects
	beego.Router("/project/new", &controllers.ProjectController{}, "get,post:NewProject")
	beego.Router("/project/:id:int/update/", &controllers.ProjectController{}, "get,post:Update")
	beego.Router("/project/:id:int/delete", &controllers.ProjectController{}, "get,post:Remove")
	beego.Router("/project/:id:int/preview", &controllers.ProjectController{}, "get,post:Preview")
	beego.Router("/project/:id:int/deploy", &controllers.ProjectController{}, "get,post:Deploy")
	beego.Router("/projects/list", &controllers.ProjectController{}, "get:List")

	// Pages
	beego.Router("/page/:projectID:int/:pageID:int/update", &controllers.PageController{}, "get,post:Update")
	beego.Router("/page/:projectID:int/:pageID:int/preview", &controllers.PageController{}, "get,post:Preview")

	// Section
	beego.Router("/section/:pageID:int/:sectionID:int/update", &controllers.SectonController{}, "get,post:Update")

	// Subsection
	beego.Router("/subsection/:sectionID:int/:subSectionID:int/update", &controllers.SubSectionController{}, "get,post:Update")

	// Lora
	beego.Router("/pricing", &controllers.LoraController{}, "get:Pricing")
	beego.Router("/services", &controllers.LoraController{}, "get:Services")
	beego.Router("/contacts", &controllers.LoraController{}, "get:Contacts")
	beego.Router("/terms", &controllers.LoraController{}, "get:Terms")
	beego.Router("/legal", &controllers.LoraController{}, "get:Legal")
	beego.Router("/help", &controllers.LoraController{}, "get:Help")

	// clearance
	cls := filters.NewBaseClearance()
	cls.Register(filters.NewUser("xshabe"), filters.LEVEL_SIX, "/")
	cls.ClearUp()
}
