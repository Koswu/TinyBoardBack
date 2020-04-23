package logging

import (
	"boarderbackend/pkgs/setting"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func getLogFilePath() string{
	return fmt.Sprintf("%s", setting.Logging.LogFilePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%v.%s",
		setting.Logging.LogFileName,
		time.Now().Format(setting.Logging.TimeFormat),
		setting.Logging.LogFileExt)
	return path.Join(prefixPath, suffixPath)

}


func openLogFile(filePath string) *os.File {
	_, err :=os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission Error: %v", err)
	}

	handle, err :=os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to Open File: %v", handle)
	}
	return handle
}

func mkDir() {
	err :=os.MkdirAll(getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}