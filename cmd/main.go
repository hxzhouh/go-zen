package main

import (
	"flag"
	"html/template"
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
	// 设置模板函数映射
	ginServer.SetFuncMap(template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"sub": func(a, b int) int { return a - b },
		"add": func(a, b int) int { return a + b },
		"unescapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	ginServer.Static("/assets", "./assets")
	ginServer.StaticFile("/favicon.ico", "./assets/icons-8-favicon-16.png")

	ginServer.LoadHTMLGlob("templates/*")
	route.Setup(env, timeout, ginServer)
	go func() {
		err := ginServer.Run(env.ServerAddress)
		if err != nil {
			log.Fatal("Server can't start: ", err)
			return
		}
	}()
	// pr := sqlite.NewPostRepository(storage.DefaultStorage)
	// posts := PostFile.ReadAllFile("./post")
	// for _, post := range posts {
	// 	pr.Create(post.ToPost())
	// }
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for i := range c {
		slog.Info("receive exit signal ", i.String(), ",exit...")
		app.Close()
		// os.Remove("./go_zen.db") //todo remove db file, test only
		os.Exit(0)
	}
}
