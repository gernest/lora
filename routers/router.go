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

	// Filters
	filters.ClearAccounts("xshabe")
}
