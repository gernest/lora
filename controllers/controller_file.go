package controllers

import (
	"github.com/astaxie/beego"
)

type FilesController struct {
	MainController
}

func (f *FilesController) Upload() {
	sess := f.ActivateContent("files/upload")
	f.LayoutSections["JScripts"] = "jscript/upload.html"
	flash := beego.NewFlash()
	if sess == nil {
		flash.Error("Need to login")
		flash.Store(&f.Controller)
		return
	}
}
