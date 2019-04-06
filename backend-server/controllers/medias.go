package controllers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"rater/backend-server/models"
	"rater/backend-server/utils"

	"github.com/astaxie/beego"
)

// MediaController definiton.
type MediaController struct {
	BaseController
}

// Upload method.
func (c *MediaController) Upload() {
	// get multiform
	form := models.MediasForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseMediasForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	beego.Debug("ParseMediasForm:", &form)

	// verify form
	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidMediasForm:", err)
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

	// save file
	path, err := saveFile(c, form);
	if err != nil {
		c.Data["json"] = models.NewErrorInfo(fmt.Sprintf("%v", err))
		c.ServeJSON()
		return
	}

	// save to database
	uploadDate := time.Now()
	media, err := models.NewMedia(form.Name, path, form.Dtype, uploadDate)
	if err != nil {
		beego.Error("NewMedia:", err)
		c.Data["json"] = models.NewErrorInfo(ErrSystem)
		c.ServeJSON()
		return
	}
	beego.Debug("NewMedia:", media)

	if code, err := media.Insert(); err != nil {
		beego.Error("InsertMedia:", err)
		if code == models.ErrDupRows {
			c.Data["json"] = models.NewErrorInfo(ErrDupFile)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}

	// return result
	c.Data["json"] = models.NewNormalInfo(fmt.Sprintf("{\"data\":{\"link\":\"%s\"}}", beego.AppConfig.String("domainname") + path))
	c.ServeJSON()
}

// Video method.
func (c *MediaController) Download() {
	if c.GetSession("user_id") == nil {
		c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
		c.ServeJSON()
		return
	}

	file := beego.AppConfig.String("apppath") + "logs/test.log"
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, file)
}

// save file
func saveFile(c *MediaController, form models.MediasForm) (string, error) {
	if !utils.CheckDtype(form.Dtype) {
		beego.Debug("Dtype error:", form.Dtype)
		return "", fmt.Errorf(ErrInputData)
	}
	files, err := c.GetFiles("file")
	if err != nil {
		beego.Debug("GetFiles:", err)
		return "", fmt.Errorf(ErrInputData)
	}

	file := files[0]
	src, err := file.Open()
	if err != nil {
		beego.Error("Open MultipartForm File:", err)
		return "", fmt.Errorf(ErrOpenFile)
	}
	defer src.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		beego.Error("Copy File to Hash:", err)
		return "", fmt.Errorf(ErrWriteFile)
	}

	hex := fmt.Sprintf("%x", hash.Sum(nil))
	path := fmt.Sprintf("%s/%s/%s%s", beego.AppConfig.String("uploadpath"), form.Dtype, hex, filepath.Ext(file.Filename))
	dst, err := os.Create(path)

	if err != nil {
		beego.Error("Create File:", err)
		return "", fmt.Errorf(ErrWriteFile)
	}
	defer dst.Close()

	src.Seek(0, 0)
	if _, err := io.Copy(dst, src); err != nil {
		beego.Error("Copy File:", err)
		return "", fmt.Errorf(ErrWriteFile)
	}

	return path, nil
}