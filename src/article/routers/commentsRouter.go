package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddArticle",
            Router: `/addArticle`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddArticleType",
            Router: `/addArticleType`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ArticleTypeDeleteById",
            Router: `/articleType/deleteById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "DeleteById",
            Router: `/deleteById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "HandleAddArticle",
            Router: `/handleAddArticle`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ShowAddType",
            Router: `/showAddType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ShowArticle",
            Router: `/showArticle`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ShowContent",
            Router: `/showContent`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ToUpdate",
            Router: `/toUpdate`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:ArticleController"] = append(beego.GlobalControllerRouter["article/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "UpdateById",
            Router: `/updateById`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:LoginController"] = append(beego.GlobalControllerRouter["article/controllers:LoginController"],
        beego.ControllerComments{
            Method: "HandleLogin",
            Router: `/handleLogin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:LoginController"] = append(beego.GlobalControllerRouter["article/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:RegController"] = append(beego.GlobalControllerRouter["article/controllers:RegController"],
        beego.ControllerComments{
            Method: "HandleReg",
            Router: `/handleReg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["article/controllers:RegController"] = append(beego.GlobalControllerRouter["article/controllers:RegController"],
        beego.ControllerComments{
            Method: "ShowReg",
            Router: `/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
