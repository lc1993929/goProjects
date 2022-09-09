package controllers

import (
	"article/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type TypeController struct {
	web.Controller
}

// @router /showAddType [get]
func (c *ArticleController) ShowAddType() {
	newOrm := orm.NewOrm()
	var articleTypes []models.ArticleType
	_, err := newOrm.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		logs.Error(err)
		c.TplName = "index.html"
	}
	c.Data["types"] = articleTypes
	c.TplName = "addType.html"
}

// @router /addArticleType [post]
func (c *ArticleController) AddArticleType() {
	typeName := c.GetString("typeName")
	articleType := models.ArticleType{TypeName: typeName}
	newOrm := orm.NewOrm()
	_, err := newOrm.Insert(&articleType)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/showAddType", 302)
}

// @router /articleType/deleteById [get]
func (c *ArticleController) ArticleTypeDeleteById() {
	id, err := c.GetInt("id")
	if err != nil {
		logs.Error(err)
	}
	articleType := models.ArticleType{Id: id}
	newOrm := orm.NewOrm()
	_, err = newOrm.Delete(&articleType)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/showAddType", 302)
}
