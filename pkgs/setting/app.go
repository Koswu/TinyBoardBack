package setting

var App struct{
	PageSize int
	JwtSecret string
}

func loadApp(){
	appCfg, err := Cfg.GetSection("app")
	if err != nil {
		sectionErr("app", err)
	}
	App.PageSize = appCfg.Key("PAGE_SIZE").MustInt(10)
	App.JwtSecret = appCfg.Key("JWT_SECRET").MustString("!@#$%^&*()")
}