package controllers

import (
	"net/http"
	"path"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils"
)

type SubdomainController struct {
	beego.Controller
}

func (s *SubdomainController) ServeSite() {
	serverStaticRouter(s.Ctx)
}

func ServeSite(ctx *context.Context) {
	serverStaticRouter(ctx)
}
func serverStaticRouter(ctx *context.Context) {
	projectPath := ctx.Input.GetData("projectPath").(string)
	staticDir := path.Join(projectPath, "www")
	if ctx.Input.Method() != "GET" && ctx.Input.Method() != "HEAD" {
		return
	}
	requestPath := path.Clean(ctx.Input.Request.URL.Path)
	if requestPath == "" {
		requestPath = "index.html"
	}
	file := path.Join(staticDir, requestPath)
	if utils.FileExists(file) {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, file)
		return
	} else {
		http.NotFound(ctx.ResponseWriter, ctx.Request)
		return
	}
}
