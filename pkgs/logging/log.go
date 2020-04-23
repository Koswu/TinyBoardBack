package logging

import (
	"boarderbackend/pkgs/setting"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	logFile            *os.File
	defaultPrefix      = ""
	defaultCallerDepth = 2

	logger *log.Logger
	levelFlags = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

type level int

const (
	levelDebug = iota
	levelInfo
	levelWarning
	levelError
	levelFatal
)

func setPrefix(nowLevel level){
	var logPrefix = ""
	_, file, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s] [%v:%v]",
			levelFlags[nowLevel], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[nowLevel])
	}
	logger.SetPrefix(logPrefix)
}

func init(){
	if setting.Logging.IsStdout {
		logFile = os.Stdout
	} else {
		filePath := getLogFileFullPath()
		logFile = openLogFile(filePath)
	}
	logger = log.New(logFile, defaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}){
	setPrefix(levelDebug)
	logger.Println(v)
}

func Info(v ...interface{}){
	setPrefix(levelInfo)
	logger.Println(v)
}

func Warn(v ...interface{}){
	setPrefix(levelWarning)
	logger.Println(v)
}

func Error(v ...interface{}){
	setPrefix(levelError)
	logger.Println(v)
}

func Fatal(v ...interface{}){
	setPrefix(levelFatal)
	logger.Println(v)
	os.Exit(1)
}
