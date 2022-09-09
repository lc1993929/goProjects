package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello/controllers:Test6Controller"] = append(beego.GlobalControllerRouter["hello/controllers:Test6Controller"],
        beego.ControllerComments{
            Method: "GetUserById",
            Router: "/test6/getUserById/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
