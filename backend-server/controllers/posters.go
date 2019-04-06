package controllers

import (
	"net/http"
	"time"

	"rater/backend-server/models"
	"rater/backend-server/utils"

	"github.com/astaxie/beego"
)

// PosterController definiton.
type PosterController struct {
	BaseController
}

// Upload method.
func (c *PosterController) Upload() {
	// get multiform
	form := models.PostersForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParsePostersForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	beego.Debug("ParsePostersForm:", &form)

	// verify form
	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidPostersForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}

	// check authentication
	if c.GetSession("user_id") != form.Name {
		c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
		c.ServeJSON()
		return
	}

	// save to database
	uploadDate := time.Now()
	poster, err := models.NewPoster(&form, uploadDate, utils.StatusUploaded)
	if err != nil {
		beego.Error("NewPoster:", err)
		c.Data["json"] = models.NewErrorInfo(ErrSystem)
		c.ServeJSON()
		return
	}
	beego.Debug("NewPoster:", poster)

	if code, err := poster.Insert(); err != nil {
		beego.Error("InsertPoster:", err)
		if code == models.ErrDupRows {
			c.Data["json"] = models.NewErrorInfo(ErrDupFile)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}

	// return result
	c.Data["json"] = models.NewNormalInfo("Succes")
	c.ServeJSON()
}

// Video method.
func (c *PosterController) Download() {
	if c.GetSession("user_id") == nil {
		c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
		c.ServeJSON()
		return
	}

	file := beego.AppConfig.String("apppath") + "logs/test.log"
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, file)
}