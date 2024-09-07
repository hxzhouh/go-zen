package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/api/route"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	_ "github.com/hxzhouh/go-zen.git/utils"
)

var configFileName string

func init() {
	flag.StringVar(&configFileName, "config", "go_zen.env", "set config file name")
	flag.Parse()
}

// @title go_zen blog system API
// @version 1.0
// @description IvanApi Service
// @BasePath /
// @Param   request_id  header  string  true  "Request ID"
// @host 127.0.0.1:8081
// @query.collection.format multi
func main() {
	slog.Info("main", "configFileName", configFileName)
	if len(configFileName) == 0 {
		slog.Error("config file name is required")
		return
	}
	app := bootstrap.App(configFileName)
	env := app.Env
	timeout := time.Duration(env.ContextTimeout) * time.Second
	ginServer := gin.Default()
	ginServer.Static("/assets", "./assets")
	ginServer.StaticFile("/favicon.ico", "./assets/icons-8-favicon-16.png")
	ginServer.LoadHTMLGlob("templates/*.html")
	route.Setup(env, timeout, ginServer)
	go func() {
		err := ginServer.Run(env.ServerAddress)
		if err != nil {
			log.Fatal("Server can't start: ", err)
			return
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for i := range c {
		slog.Info("receive exit signal ", i.String(), ",exit...")
		app.Close()
		os.Exit(0)
	}
}
