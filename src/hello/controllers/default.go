package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Test1Controller struct {
	beego.Controller
}

func (c *Test1Controller) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "hahaha"
	c.TplName = "test1.html"
}

func (c *Test1Controller) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "blublublu"
	c.TplName = "test1.html"
}

type Test2Controller struct {
	beego.Controller
}

func (c *Test2Controller) Test2() {
	fmt.Println("test2")
	c.TplName = "index.tpl"
}

type Test3Controller struct {
	beego.Controller
}

func (c *Test3Controller) Test31() {
	fmt.Println("test31")
	c.TplName = "index.tpl"
}

func (c *Test3Controller) Test32() {
	fmt.Println("test32")
	c.TplName = "index.tpl"
}

type Test4Controller struct {
	beego.Controller
}

func (u *Test4Controller) HelloWorld() {
	fmt.Println("test4")
	u.Ctx.WriteString("hello, world")
}

type Test6Controller struct {
	beego.Controller
}

// GetUserById @router /test6/getUserById/:id [get]
func (u *Test6Controller) GetUserById() {
	fmt.Println(u.GetString(":id" + "controller"))
	//fmt.Println(ctx.Request.URL.Path + "context")
	u.Ctx.WriteString("test6")
}

func FunTest(ctx *context.Context) {
	//ctx.Request.FormFile("testFile")
	ctx.Redirect(404, "/")
	ctx.WriteString("test7")
}
