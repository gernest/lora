package main

import (
	"path/filepath"

	"github.com/astaxie/beego"

	"github.com/gernest/lora/models"
	"github.com/gernest/lora/utilities/logs"

	_ "github.com/gernest/lora/models"
	_ "github.com/gernest/lora/routers"
)

func init() {

	// Set Logging
	beego.SetLogFuncCall(true)

	// Make sure index.htm shows up in project preview
	beego.DirectoryIndex = true

	beego.SessionOn = true
	beego.SessionName = "xshabe"

	// make sure we will preview our shitty sites
	loadProjectPreview()
}

var logThis = logs.NewLoraLog()

func main() {
	models.RunMigrations()

	logThis.Info("========================")
	logThis.Info("******HAPY LORA*********")
	logThis.Info("========================")

	beego.Run()
}

func loadProjectPreview() {
	db, err := models.Conn()
	if err != nil {
		logThis.Info("some fish %v", err)
	}
	ps := []models.Project{}
	err = db.Select("name, project_path").Find(&ps).Error
	if err != nil {
		logThis.Info("Proublem updating previw paths ** %v **", err)
	}
	for _, v := range ps {
		staticPath := filepath.Join(v.ProjectPath, "public")
		previewPath := "/preview/" + v.Name
		beego.SetStaticPath(previewPath, staticPath)
	}
}
