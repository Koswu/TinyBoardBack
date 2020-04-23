package models

import (
	"boarderbackend/pkgs/logging"
	"boarderbackend/pkgs/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"path"
	"time"
)

var db *gorm.DB

type Model struct {
	ID int`gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope){
	scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (model *Model) BeforeUpdate(scope gorm.Scope){
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}


func init() {
	var (
		err error
		tablePrefix string
	)
	tablePrefix = setting.DB.TablePrefix
	switch setting.DB.Type{
	case "sqlite":
		{
			dir, _ := path.Split(setting.DB.Path)
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				logging.Fatal("Mkdir failed", err)
			}
			db, err = gorm.Open("sqlite3", setting.DB.Path)
			if err != nil {
				logging.Fatal("Open sqlite dataBase failed", err, setting.DB.Path)
			}
		}
	case "mysql":
		{
			db, err = gorm.Open("mysql",
				fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
					setting.DB.User,
					setting.DB.PassWord,
					setting.DB.Host,
					setting.DB.Name))
			if err != nil {
				logging.Fatal("Connect to database failed", err)
			}
		}
	default:
		logging.Fatal("Not Support Database type %s", setting.DB.Type)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	createTables()
}


func CloseDB() {
	defer db.Close()
}

