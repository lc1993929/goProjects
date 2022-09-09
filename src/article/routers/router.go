package routers

import (
	"article/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter("/article/*", beego.BeforeRouter, controllers.LoginFilter)
	beego.Router("/", &controllers.LoginController{})
	beego.Include(&controllers.RegController{})
	beego.Include(&controllers.LoginController{})
	beego.Include(&controllers.ArticleController{})
	beego.Include(&controllers.TypeController{})
	//beego.Post("/handleReg", controllers.HandleReg)

	_ = beego.AddFuncMap("nextPage", controllers.NextPage)
	_ = beego.AddFuncMap("prePage", controllers.PrePage)
}
