package controllers

import (
	"github.com/astaxie/beego"
)

// MainController definition.
type MainController struct {
	beego.Controller
}

// Get method.
func (c *MainController) Get() {
	c.Data["Website"] = "rater.me"
	c.Data["Email"] = "wenchuang.liu@hotmail.com"
	c.TplName = "index.tpl"
	c.Render()
}
