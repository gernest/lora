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

package imgr

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/gernest/lora/utils/logs"
)

var logThis = logs.NewLoraLog()

type ImageManager struct {
	Source      string
	Destinalion string
	Images      []*Image
	AllowedExt  []string
}

type Image struct {
	Name       string
	Dimensions []int
	Path       string
	Ext        string
}

func (i *ImageManager) AddImage(path string) error {

	img, err := getImageDetails(path)
	if err != nil {
		return err
	}
	pic := newImage(img, path)
	n := append(i.Images, pic)
	i.Images = n
	return nil
}
func (i *ImageManager) LoadFromSource() error {
	return i.loadImagesFromSource()
}
func (i *ImageManager) loadImagesFromSource() error {
	err := filepath.Walk(i.Source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		match := filterByExtension(path, i.AllowedExt)
		if match {
			err = i.AddImage(path)
			return err
		}
		return nil
	})
	return err
}

func filterByExtension(path string, exts []string) bool {

	xt := filepath.Ext(path)
	for _, v := range exts {
		if xt == v {
			return true
		}
	}
	return false
}

func getNameWithoutExt(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func getImageDetails(path string) (image.Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return image.Config{}, err
	}
	img, _, err := image.DecodeConfig(file)
	return img, err
}
func newImage(img image.Config, path string) *Image {
	return &Image{
		Name:       getNameWithoutExt(path),
		Dimensions: []int{img.Height, img.Width},
		Path:       path,
		Ext:        filepath.Ext(path),
	}

}
func NewImageManager(src string, dest string, allow []string) *ImageManager {
	return &ImageManager{
		Source:      src,
		Destinalion: dest,
		Images:      []*Image{},
		AllowedExt:  allow,
	}
}
