package filters

import (
	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/gernest/lora/controllers"
	"github.com/gernest/lora/models"
)

func StaticProxy(ctx *context.Context) {
	db, err := models.Conn()
	if err != nil {
		return
	}
	defer db.Close()
	projects := []models.Project{}
	db.Debug().Find(&projects)

	if ctx.Input.SubDomains() != "" {
		for _, project := range projects {
			if ctx.Input.SubDomains() == project.Name {
				ctx.Input.SetData("projectPath", project.ProjectPath)
				controllers.ServeSite(ctx)
				break
			}
		}
		if ctx.Input.GetData("projectPath") == nil {
			http.NotFound(ctx.ResponseWriter, ctx.Request)
			return
		}
	}

}
