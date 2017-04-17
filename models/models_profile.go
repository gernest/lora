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

package models

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"path/filepath"

	"github.com/astaxie/beego"

	"github.com/1l0/identicon"
	"github.com/kardianos/osext"
)

func (p *Profile) GenerateIdenticon(base, s string) error {
	var (
		gPath           string
		identiconPath   string
		defaultIconPath string
		link            string
	)

	id := identicon.New()
	identiconPath = beego.AppConfig.String("identiconsPath")
	defaultIconPath = "static/img/identicons"
	baseDir, _ := osext.ExecutableFolder()

	name := getHashName(s)

	gPath = filepath.Join(baseDir, identiconPath)
	link = fmt.Sprintf("/%s/%s.png", identiconPath, name)

	if identiconPath == "" {
		gPath = filepath.Join(base, defaultIconPath)
		link = fmt.Sprintf("/%s/%s.png", defaultIconPath, name)
	}

	saveTo := filepath.Join(gPath, name)

	if err := generateProfileImg(id, saveTo); err != nil {
		return err
	}

	p.Photo = link
	return nil
}

func getHashName(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func generateProfileImg(i *identicon.Identicon, p string) error {
	n := rand.Intn(5)
	t := []int{identicon.Normal, identicon.Mirrorh, identicon.Mirrorv}
	switch n {
	case 0:
		i.Theme = identicon.White
		i.Type = t[rand.Intn(len(t))]
	case 1:
		i.Theme = identicon.White
		i.Type = t[rand.Intn(len(t))]

	case 2:
		i.Theme = identicon.Gray
		i.Type = t[rand.Intn(len(t))]
	default:
		i.Theme = identicon.Free
		i.Type = identicon.Mirrorv
	}
	return i.GeneratePNGToFile(p)
}
