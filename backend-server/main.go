package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/astaxie/beego/plugins/cors"

	"rater/backend-server/controllers"
	_ "rater/backend-server/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func handleSignals(c chan os.Signal) {
	switch <-c {
	case syscall.SIGINT, syscall.SIGTERM:
		fmt.Println("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		fmt.Println("Shutdown gracefully, bye...")
		// do graceful shutdown
	}

	os.Exit(0)
}

func main() {
	graceful, _ := beego.AppConfig.Bool("graceful")
	if !graceful {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go handleSignals(sigs)
	}

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	mode := beego.AppConfig.String("runmode")
	if mode == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}

	beego.ErrorController(&controllers.ErrorController{})


	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"https://*.foo.com", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Run()
}
