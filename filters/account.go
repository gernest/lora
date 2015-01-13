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

package filters

import (
	"github.com/gernest/lora/utils/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var logThis = logs.NewLoraLog()

type ResourceClearance interface {
	Clear()
}

type ClearanceObject interface {
	ClearanceLevel() int
}

type BaseClearance struct {
	Ctx     *context.Context
	Objects []*baseClearanceObject
}
type baseClearanceObject struct {
	object ClearanceObject
	level  int
	pos    int
	route  string
}

func (o *baseClearanceObject) Clear() {
	beego.InsertFilter(o.route, o.pos, func(ctx *context.Context) {
		if o.object.ClearanceLevel() < o.level {
			logThis.Info("No Permission")
		}
		logThis.Info("Permitted")
	})
}

func (b *BaseClearance) Register(o ClearanceObject, level int, route string) {
	base := &baseClearanceObject{
		object: o,
		level:  level,
		pos:    beego.BeforeRouter,
		route:  route,
	}
	objects := append(b.Objects, base)
	b.Objects = objects
}

func (b *BaseClearance) ClearUp() {
	for _, ob := range b.Objects {
		ob.Clear()
	}
}
