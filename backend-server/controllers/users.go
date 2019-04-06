package controllers

import (
	"time"

	"rater/backend-server/models"

	"github.com/astaxie/beego"
)

// UserController definiton.
type UserController struct {
	BaseController
}

// Register method.
func (c *UserController) Register() {
	form := models.RegisterForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseRegsiterForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	beego.Debug("ParseRegsiterForm:", &form)

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidRegsiterForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}

	regDate := time.Now()
	user, err := models.NewUser(&form, regDate)
	if err != nil {
		beego.Error("NewUser:", err)
		c.Data["json"] = models.NewErrorInfo(ErrSystem)
		c.ServeJSON()
		return
	}
	beego.Debug("NewUser:", user)

	if code, err := user.Insert(); err != nil {
		beego.Error("InsertUser:", err)
		if code == models.ErrDupRows {
			c.Data["json"] = models.NewErrorInfo(ErrDupUser)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = models.NewNormalInfo("Succes")
	c.ServeJSON()
}

// Login method.
func (c *UserController) Login() {
	form := models.LoginForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	beego.Debug("ParseLoginForm:", &form)

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidLoginForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}

	user := models.User{}
	if code, err := user.FindByName(form.Name); err != nil {
		beego.Error("FindUserByName:", err)
		if code == models.ErrNotFound {
			c.Data["json"] = models.NewErrorInfo(ErrNoUser)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}
	beego.Debug("UserInfo:", &user)

	if ok, err := user.CheckPass(form.Password); err != nil {
		beego.Error("CheckUserPass:", err)
		c.Data["json"] = models.NewErrorInfo(ErrSystem)
		c.ServeJSON()
		return
	} else if !ok {
		c.Data["json"] = models.NewErrorInfo(ErrPass)
		c.ServeJSON()
		return
	}
	user.ClearPass()

	c.SetSession("user_id", form.Name)

	c.Data["json"] = &models.LoginInfo{Code: 0, UserInfo: &user}
	c.ServeJSON()
}

// Logout method.
func (c *UserController) Logout() {
	form := models.LogoutForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLogoutForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	beego.Debug("ParseLogoutForm:", &form)

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidLogoutForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}

	if c.GetSession("user_id") != form.Name {
		c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
		c.ServeJSON()
		return
	}

	c.DelSession("user_id")

	c.Data["json"] = models.NewNormalInfo("Succes")
	c.ServeJSON()
}

// Passwd method.
func (c *UserController) Passwd() {
	form := models.PasswdForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParsePasswdForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	beego.Debug("ParsePasswdForm:", &form)

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidPasswdForm:", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}

	if c.GetSession("user_id") != form.Name {
		c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
		c.ServeJSON()
		return
	}

	code, err := models.ChangePass(form.Name, form.OldPass, form.NewPass)
	if err != nil {
		beego.Error("ChangeUserPass:", err)
		if code == models.ErrNotFound {
			c.Data["json"] = models.NewErrorInfo(ErrNoUserPass)
		} else if code == models.ErrDatabase {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrSystem)
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = models.NewNormalInfo("Succes")
	c.ServeJSON()
}
