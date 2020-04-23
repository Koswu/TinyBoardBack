package setting

var DB struct{
	Type string
	Name string
	Path string
	TablePrefix string
	User string
	PassWord string
	Host string
}

func loadDB(){
	dbSection:= getSection("database")
	DB.Type = dbSection.Key("TYPE").MustString("sqlite")
	DB.TablePrefix = dbSection.Key("TABLE_PREFIX").MustString("board")
	if DB.Type != "sqlite"{
		connectSection := getSection("database.connect")
		DB.User = connectSection.Key("USER").MustString("root")
		DB.PassWord = connectSection.Key("PASSWORD").MustString("root")
		DB.Host = connectSection.Key("HOST").MustString("localhost:3306")
		DB.Name = connectSection.Key("NAME").MustString("board")
	} else {
		fileSection := getSection("database.file")
		DB.Path = fileSection.Key("DB_PATH").MustString("runtime/app.db")
	}
}