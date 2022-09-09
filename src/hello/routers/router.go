package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test1", &controllers.Test1Controller{})
	beego.Router("/test2", &controllers.Test2Controller{}, "get,post:Test2")
	beego.Router("/test3", &controllers.Test3Controller{}, "get:Test31;post:Test32")
	beego.AutoRouter(&controllers.Test4Controller{}).AutoPrefix("pre", &controllers.Test4Controller{})
	beego.Get("/test5/user/:id", func(ctx *context.Context) {
		ctx.WriteString("test5" + ctx.Request.URL.Path + ctx.Input.Param(":id"))
	})
	beego.Include(&controllers.Test6Controller{})
	beego.Get("/test7/funcTest", controllers.FunTest)
	//beego.Router("/test3", &controllers.Test2Controller{}, "*:Test31")

	beego.AutoRouter(&controllers.PetController{})

}
