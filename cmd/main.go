package main

import (
	"fast-go/conf"
	"fast-go/internal/app/config"
	"fast-go/internal/app/models"
	"fast-go/internal/app/routers"
	"fast-go/pkg/gredis"
	"fast-go/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	//setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	config.LoadGlobalConfig(configFile)
}

const (
	configFile = "../conf/config.toml"
)

func main() {
	gin.SetMode(conf.App.Server.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := conf.App.Server.ReadTimeout * time.Second
	writeTimeout := conf.App.Server.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", conf.App.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
