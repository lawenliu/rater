package routers

import (
	"rater/backend-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),
			beego.NSRouter("/passwd", &controllers.UserController{}, "post:Passwd"),
		),
		beego.NSNamespace("/medias",
			beego.NSRouter("/upload", &controllers.MediaController{}, "post:Upload"),
			beego.NSRouter("/download", &controllers.MediaController{}, "post:Download"),
		),
		beego.NSNamespace("/posters",
			beego.NSRouter("/upload", &controllers.PosterController{}, "post:Upload"),
			beego.NSRouter("/download", &controllers.PosterController{}, "post:Download"),
		),

	)
	beego.AddNamespace(ns)
}
