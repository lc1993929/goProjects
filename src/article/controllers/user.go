package controllers

import (
	"article/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"time"
)

type RegController struct {
	web.Controller
}

// @router /register [get]
func (c *RegController) ShowReg() {
	c.TplName = "register.html"
}

// @router /handleReg [post]
func (c *RegController) HandleReg() {
	userName := c.GetString("userName")
	password := c.GetString("password")
	logs.Info(userName, password)
	valid := validation.Validation{}
	valid.Required(userName, "userName").Message("userName不能为空")
	valid.Required(password, "password").Message("password不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
		}
	}

	newOrm := orm.NewOrm()
	user := models.User{UserName: userName, Password: password}
	_, err := newOrm.Insert(&user)
	if err != nil {
		logs.Error(err)
	}

	c.TplName = "login.html"
}

type LoginController struct {
	web.Controller
}

func (c *LoginController) Get() {
	userName := c.Ctx.GetCookie("userName")
	c.Data["userName"] = userName
	c.TplName = "login.html"
}

// @router /handleLogin [post]
func (c *LoginController) HandleLogin() {
	userName := c.GetString("userName")
	password := c.GetString("password")
	logs.Info(userName, password)
	valid := validation.Validation{}
	valid.Required(userName, "userName").Message("userName不能为空")
	valid.Required(password, "password").Message("password不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
		}
		c.TplName = "login.html"
		return
	}

	user := models.User{}
	newOrm := orm.NewOrm()
	userTable := newOrm.QueryTable(&user)
	err := userTable.Filter("userName", userName).One(&user)
	if err != nil {
		logs.Error(err)
		if err == orm.ErrNoRows {
			logs.Warn("未找到用户名%s", userName)
			c.TplName = "login.html"
			return
		}
	}
	if user.Password != password {
		logs.Warn("密码错误")
		c.TplName = "login.html"
		return
	}
	remember := c.GetString("remember")
	if remember == "on" {
		c.Ctx.SetCookie("userName", userName, time.Second*3600)
	} else {
		c.Ctx.SetCookie("userName", userName, -1)
	}
	err = c.SetSession("userName", userName)
	if err != nil {
		logs.Error(err)
		c.TplName = "login.html"
		return
	}

	c.Redirect("/showArticle", 302)
}

// @router /logout [get]
func (c *LoginController) Logout() {
	err := c.DelSession("userName")
	if err != nil {
		logs.Error(err)
	}
	c.TplName = "login.html"
}

/*func HandleReg(ctx *context.Context) {
	userName := ctx.Request.PostFormValue("userName")
	password := ctx.Request.PostFormValue("password")
	logs.Info(userName, password)
	ctx.Output.("success")
}*/

func LoginFilter(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302, "/")
	}
}
