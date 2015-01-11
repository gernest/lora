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

package main

import (
	"path/filepath"

	"github.com/astaxie/beego"

	"github.com/gernest/lora/models"
	"github.com/gernest/lora/utils/logs"

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
	loadDeployedApps()
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
	defer db.Close()
	if err != nil {
		logThis.Info("some fish %v", err)
	}
	ps := []models.Project{}
	err = db.Select("name, project_path").Find(&ps).Error
	if err != nil {
		logThis.Info("Proublem updating previw paths ** %v **", err)
	}
	for _, v := range ps {
		staticPath := filepath.Join(v.ProjectPath, "www")
		previewPath := "/preview/" + v.Name
		beego.SetStaticPath(previewPath, staticPath)
	}
}

func loadDeployedApps() {
	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		logThis.Info("some fish %v", err)
	}
	ps := []models.Project{}
	err = db.Select("name, project_path").Find(&ps).Error
	if err != nil {
		logThis.Info("Proublem updating previw paths ** %v **", err)
	}
	for _, v := range ps {
		staticPath := filepath.Join(v.ProjectPath, "www")
		previewPath := "/apps/" + v.Name
		beego.SetStaticPath(previewPath, staticPath)
	}
}
