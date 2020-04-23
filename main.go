package main

import (
	"boarderbackend/models"
	_ "boarderbackend/pkgs/logging"
	"boarderbackend/pkgs/setting"
	"boarderbackend/router"
	"fmt"
	"net/http"
	"time"
)

func main() {
	defer models.CloseDB()
	router := router.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.Server.HttpPort),
		Handler:router,
		ReadTimeout: time.Duration(setting.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(setting.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1<<20,
	}
	_ = s.ListenAndServe()
}