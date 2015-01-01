package models

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/gernest/lora/utilities/logs"

	"github.com/astaxie/beego"
	cp "github.com/gernest/lora/utilities/copy"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var logThis = logs.NewLoraLog()

func connStr() string {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_passwd")
	db_name := beego.AppConfig.String("db_name")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	dns = fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s", db_name, db_host, db_port, db_user, db_pass, db_sslmode)
	return dns
}

// Conn provied a *gorm.DB which is ready for interacting with database
func Conn() (*gorm.DB, error) {
	var dbPath, dialect string
	var err error
	var db gorm.DB
	dbPath = connStr()
	dialect = beego.AppConfig.String("db_dialect")
	db, err = gorm.Open(dialect, dbPath)
	//db.LogMode(true)
	return &db, err
}

func RunMigrations() {
	logThis.Info("Running Migrations")
	db, err := Conn()
	if err != nil {
		beego.Info("Failed to initialize database with error ", err)
		return
	}
	//db.LogMode(true)
	//db.DropTableIfExists(new(Account))
	db.DropTableIfExists(new(Project))
	db.DropTableIfExists(new(Page))
	db.DropTableIfExists(new(Section))
	db.DropTableIfExists(new(SubSection))
	//db.DropTableIfExists(new(Profile))
	db.AutoMigrate(new(Account), new(Profile), new(Project), new(Page), new(Section), new(SubSection))
	logThis.Success("migrations succeded")
}

func NewLoraProject(base string, name string, template string, theme string) (Project, error) {
	p := new(Project)
	err := p.Initialize(base, name, template, theme)
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}

	err = p.GenScaffold()
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}
	err = copyTheme(p)
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}
	err = p.InitDir()
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}
	err = p.LoadConfigFile()
	if err != nil {
		logThis.Debug("Trouble to create new project *%v*", err)
		return *p, err
	}

	p.GenLorem() // fill content with dummy data

	return *p, nil
}

func copyTheme(p *Project) error {
	themeDir := beego.AppConfig.String("themesDir")
	if themeDir == "" {
		themeDir = "themes"
	}

	sourceDir := filepath.Join(filepath.Join(p.BaseDir, themeDir), p.Theme)
	destDir := filepath.Join(filepath.Join(p.ProjectPath, "themes"), p.Theme)

	err := cp.CopyDir(sourceDir, destDir)
	if err != nil {
		logThis.Debug("Trouble copying theme *%v*", err)
		return err
	}
	return nil
}

func sanitizeHTMLField(s string) []byte {
	unsafe := blackfriday.MarkdownCommon([]byte(s))
	safe := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return safe
}

func getFuncCall(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		_, filename := path.Split(file)
		s := fmt.Sprintf("[%s : %d] ", filename, line)
		return s
	}
	return ""
}
