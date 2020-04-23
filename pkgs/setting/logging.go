package setting


var Logging struct{
	IsStdout bool
	TimeFormat string
	LogFilePath string
	LogFileName string
	LogFileExt string
}

func loadLogging(){
	logSection := getSection("logging")
	Logging.IsStdout = logSection.Key("IS_STDOUT").MustBool(true)
	Logging.TimeFormat = logSection.Key("TIME_FORMAT").MustString("2006-01-02")
	if !Logging.IsStdout {
		fileSection :=  getSection("logging.logfile")
		Logging.LogFilePath = fileSection.Key("LOG_FILE_PATH").MustString("runtime/logs/")
		Logging.LogFileName = fileSection.Key("LOG_FILE_NAME").MustString("log")
		Logging.LogFileExt = fileSection.Key("LOG_FILE_EXT").MustString("log")
	}
}