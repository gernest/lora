package filters

import (
	"errors"

	"github.com/astaxie/beego/context"
	"github.com/gernest/lora/models"
)

type User struct {
	account  *models.Account
	ctx      *context.Context
	sessName string
}

func (u *User) ClearanceLevel() (int, error) {
	err := u.loadAccount()
	if err != nil {
		logThis.Info("Fuck thhe user is not there, failed t %v", err)
		return 0, err
	}
	return u.account.ClearanceLevel, nil
}
func (u *User) NewContext(ctx *context.Context) {
	u.ctx = ctx
}

func (u *User) loadAccount() error {
	sess := u.ctx.Input.Session(u.sessName)

	if sess == nil {
		return errors.New("Session not found")
	}
	m := sess.(map[string]interface{})

	db, err := models.Conn()
	defer db.Close()
	if err != nil {
		return err
	}
	a := models.Account{}
	a.Email = m["email"].(string)
	query := db.Where("email= ?", a.Email).First(&a)
	if query.Error != nil {
		return err
	}
	u.account = &a
	return nil
}

func NewUser(sess string) *User {
	return &User{
		sessName: sess,
	}
}
