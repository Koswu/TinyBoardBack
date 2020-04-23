package setting

import "log"

func loadIniError(err error){
	log.Fatalf("[FATAL] failed to parse 'conf/app.ini': %v", err)
}

func sectionErr(sectionName string, err error){
	log.Fatalf("Load section '%s' failed %v", sectionName, err)
}
