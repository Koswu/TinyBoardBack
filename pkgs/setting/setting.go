package setting

import (
	"gopkg.in/ini.v1"
)

var Cfg *ini.File
var IsDebug bool

func init(){
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		loadIniError(err)
	}
	loadBase()
	loadLogging()
	loadApp()
	loadDB()
	loadServer()
}

func loadBase(){
	IsDebug = Cfg.Section("").Key("IS_DEBUG").MustBool()
}

func getSection(sec string)*ini.Section{
	section, err := Cfg.GetSection(sec)
	if err != nil {
		sectionErr(sec, err)
	}
	return section
}
