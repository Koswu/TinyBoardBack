package models

import (
	"boarderbackend/pkgs/logging"
	"reflect"
)

func createTables() {
	createTable(&Comment{})
	createTable(&Auth{})
}

func createTable(typo interface{}) {
	tp := reflect.TypeOf(typo).Elem().Name()
	if !db.HasTable(typo) {
		db.CreateTable(typo)
		logging.Warn(tp + " Table is not exist, Created")
	}
}
