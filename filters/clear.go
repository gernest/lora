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
	"github.com/gernest/lora/models"
	"github.com/gernest/lora/utils/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	LEVEL_GROUND = iota
	LEVEL_ONE
	LEVEL_TWO
	LEVEL_THREE
	LEVEL_FOUR
	LEVEL_FIVE
	LEVEL_SIX
)

var logThis = logs.NewLoraLog()

func AuthLevelSix(ctx *context.Context) {
	clearByLevel(LEVEL_SIX, ctx)
}

func AuthLevelFive(ctx *context.Context) {
	clearByLevel(LEVEL_FIVE, ctx)
}

func AuthLevelFour(ctx *context.Context) {
	clearByLevel(LEVEL_FOUR, ctx)
}

func AuthLevelThree(ctx *context.Context) {
	clearByLevel(LEVEL_THREE, ctx)
}

func AuthLevelTwo(ctx *context.Context) {
	clearByLevel(LEVEL_TWO, ctx)
}

func AuthLevelOne(ctx *context.Context) {
	clearByLevel(LEVEL_ONE, ctx)
}

func clearByLevel(perm int, ctx *context.Context) {
	sessName := beego.AppConfig.String("sessionname")
	usr := &User{}
	usr.ctx = ctx
	usr.sessName = sessName

	level, err := usr.ClearanceLevel()
	if err != nil {
		ctx.Input.SetData("user", &models.Account{})
		return
	}
	if level >= perm {
		ctx.Input.SetData("user", usr.account)
		return
	}
}
