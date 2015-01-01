package filters

import (
	"errors"
	"reflect"

	"github.com/kr/pretty"

	"github.com/astaxie/beego"

	"github.com/gernest/lora/models"

	"github.com/astaxie/beego/context"
)

type Clearance struct {
	Pattern  string
	Position int
	Level    int
	Obj      interface{}
	Ctx      *context.Context
	Session  string
}

func (c Clearance) SetPattern(s string) Clearance {
	c.Pattern = s
	return c
}
func (c Clearance) SetPosition(p int) Clearance {
	c.Position = p
	return c
}
func (c Clearance) SetLevel(level int) Clearance {
	c.Level = level
	return c
}
func (c Clearance) Clear() {
	beego.InsertFilter(c.Pattern, c.Position, func(ctx *context.Context) {
		obj, err := c.GetObjectFromSession(ctx)
		if err != nil {
			pretty.Println(err)
			return
		}
		objLevel := obj.ObjectClearanceLevel()
		if objLevel < c.Level {
			// This is not cleared so enforce the rules here
			return
		}

	})
}
func (c Clearance) SetSession(s string) Clearance {
	c.Session = s
	return c
}

func (c Clearance) ObjectClearanceLevel() int {
	objLevel := reflect.ValueOf(c.Obj).FieldByName("ClearanceLevel").Int()
	return int(objLevel)
}
func (c Clearance) GetObjectFromSession(ctx *context.Context) (Clearance, error) {
	sess := ctx.Input.Session(c.Session)

	if sess == nil {
		return c, errors.New("Session not found")
	}
	m := sess.(map[string]interface{})

	db, err := models.Conn()
	if err != nil {
		return c, err
	}
	a := models.Account{}
	a.Email = m["email"].(string)
	query := db.Where("email= ?", a.Email).First(&a)
	if query.Error != nil {
		return c, err
	}
	c.Obj = a
	c.Ctx = ctx
	return c, err
}

func New() Clearance {
	return Clearance{}
}
func ClearAccounts(sess string) {
	c := New()
	c.SetSession(sess).
		SetLevel(6).SetPosition(beego.BeforeRouter).
		SetPattern("/").Clear()
}
