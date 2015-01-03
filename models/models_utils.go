package models

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"

	"bitbucket.org/kardianos/osext"

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
	err = copyTheme(p, p.Theme)
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

func GetAvailableThemes(base string) ([]string, error) {
	list, err := getResourceList("themes", base)
	return list, err
}
func GetAvailableTemplates(base string) ([]string, error) {
	list, err := getResourceList("templates", base)
	return list, err
}

func copyTheme(p *Project, name string) error {
	themeDir := beego.AppConfig.String("themesDir")
	if themeDir == "" {
		themeDir = "themes"
	}
	if name == "" || p.Theme == "" {
		name = "loraina"
	}
	sourceDir := filepath.Join(filepath.Join(p.BaseDir, themeDir), name)
	destDir := filepath.Join(filepath.Join(p.ProjectPath, "themes"), name)

	err := cp.CopyDir(sourceDir, destDir)
	if err != nil {
		logThis.Debug("Trouble copying theme *%v*", err)
		return err
	}
	p.Theme = name
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

func getResourceList(s, location string) ([]string, error) {
	var (
		resource, resourcePath, resourceDir, base string
		resourceList                              []string
	)
	base = location
	if location == "" {
		baseD, err := osext.ExecutableFolder()
		if err != nil {
			return []string{}, err
		}
		base = baseD
	}

	switch s {
	case "themes":
		resource = "themes"
		resourceDir = beego.AppConfig.String("themesDir")
		if resourceDir == "" {
			resourceDir = resource
		}
	case "templates":
		resource = "templates"
		resourceDir = beego.AppConfig.String("templatesDIr")
		if resourceDir == "" {
			resourceDir = resource
		}
	}
	resourcePath = filepath.Join(base, resourceDir)
	if filepath.IsAbs(resourceDir) {
		resourcePath = resourceDir
	}
	file, err := os.Open(resourcePath)
	if err != nil {
		return []string{}, err
	}
	resourceList, err = file.Readdirnames(2)

	if err != nil {
		return []string{}, err
	}
	sort.Sort(sort.StringSlice(resourceList)) //Sort the list
	return resourceList, nil
}
