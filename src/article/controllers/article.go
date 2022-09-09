package controllers

import (
	"article/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/sineycoder/go-bigger/bigger"
	"github.com/sineycoder/go-bigger/types"
	"mime/multipart"
	"net/http"
	"path"
	"reflect"
	"strconv"
	"time"
)

type ArticleController struct {
	web.Controller
}

// @router /addArticle [get]
func (c *ArticleController) AddArticle() {
	var articleTypes []models.ArticleType
	newOrm := orm.NewOrm()
	_, err := newOrm.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		logs.Error(err)
	}
	c.Data["types"] = articleTypes
	c.TplName = "add.html"
}

// @router /handleAddArticle [post]
func (c *ArticleController) HandleAddArticle() {
	title := c.GetString("title")
	typeString := c.GetString("type")
	content := c.GetString("content")
	file, header, err := c.GetFile("img")

	if err != nil {
		logs.Error(err)
		if err == http.ErrMissingFile {
			c.Redirect("/addArticle", 302)
			return
		}
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			logs.Error(err)
		}
	}(file)

	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		logs.Error("上传文件格式错误")
		return
	}
	if header.Size > 500000000000 {
		logs.Error("文件过大")
		return
	}
	timeFormat := time.Now().Format("2006-01-02 15-04-05")
	imgPath := "static/img/" + timeFormat + ext
	err = c.SaveToFile("img", imgPath)
	if err != nil {
		logs.Error(err)
		c.Redirect("/addArticle", 302)
		return
	}

	o := orm.NewOrm()
	articleType := models.ArticleType{}
	err = o.QueryTable("ArticleType").Filter("TypeName", typeString).One(&articleType)
	if err != nil {
		logs.Error(err)
		c.Redirect("/addArticle", 302)
		return
	}
	article := models.Article{Title: title, Img: imgPath, Type: &articleType, Content: content}
	valid := validation.Validation{}
	b, err := valid.Valid(&article)
	if err != nil {
		logs.Error(err)
		c.Redirect("/addArticle", 302)
		return
	}
	if !b {
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
		}
		c.Redirect("/addArticle", 302)
		return
	}

	_, err = o.Insert(&article)
	if err != nil {
		logs.Error(err)
		c.Redirect("/addArticle", 302)
		return
	}

	c.Redirect("/showArticle", 302)

}

// @router /showArticle [get]
func (c *ArticleController) ShowArticle() {
	userName := c.GetSession("userName")
	if userName == nil {
		c.Redirect("/", 302)
		return
	}

	pageSize := 3

	page, err := c.GetInt("page")
	if err != nil {
		if reflect.TypeOf(err).String() == "*strconv.NumError" {
			page = 1
		} else {
			logs.Error(err)
			return
		}
	}
	typeName := c.GetString("type")
	o := orm.NewOrm()
	querySeter := o.QueryTable("article").RelatedSel("Type")
	if typeName == "" {
		articleType := models.ArticleType{}
		err := o.QueryTable("ArticleType").Limit(1).One(&articleType)
		if err != nil {
			logs.Error(err)
		}
		typeName = articleType.TypeName
	}
	querySeter = querySeter.Filter("Type__TypeName", typeName)
	c.Data["chooseType"] = typeName
	var articles []models.Article
	_, err = querySeter.Limit(pageSize, (page-1)*pageSize).All(&articles)
	if err != nil {
		logs.Error(err)
		return
	}
	c.Data["articles"] = articles

	count, err := querySeter.Count()
	if err != nil {
		logs.Error(err)
	}
	c.Data["count"] = count
	c.Data["page"] = page
	pageCount, _ := strconv.Atoi(bigger.BigDecimalValueOf(types.Long(count)).Divide(bigger.BigDecimalValueOf(types.Long(pageSize)), 0, bigger.ROUND_UP).String())
	c.Data["pageCount"] = pageCount
	c.Data["firstPage"] = page == 1
	c.Data["endPage"] = page == pageCount

	var articleTypes []models.ArticleType
	_, err = o.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		logs.Error(err)
	}
	c.Data["types"] = articleTypes

	c.TplName = "index.html"

}

// @router /showContent [get]
func (c *ArticleController) ShowContent() {
	id, err := c.GetInt("id")
	if err != nil {
		logs.Error(err)
		return
	}
	article := models.Article{Id: id}
	newOrm := orm.NewOrm()
	err = newOrm.Read(&article)
	if err != nil {
		logs.Error(err)
		return
	}
	c.Data["article"] = article
	c.TplName = "content.html"

}

// @router /toUpdate [get]
func (c *ArticleController) ToUpdate() {
	id, err := c.GetInt("id")
	if err != nil {
		logs.Error(err)
		return
	}
	article := models.Article{Id: id}
	o := orm.NewOrm()
	err = o.QueryTable("article").RelatedSel("Type").Filter("id", id).One(&article)
	if err != nil {
		logs.Error(err)
	}
	c.Data["article"] = article

	var articleTypes []models.ArticleType
	_, err = o.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		logs.Error(err)
	}
	c.Data["types"] = articleTypes
	c.TplName = "update.html"

}

// @router /deleteById [get]
func (c *ArticleController) DeleteById() {
	id, err := c.GetInt("id")
	if err != nil {
		logs.Error(err)
	}
	article := models.Article{Id: id}
	newOrm := orm.NewOrm()
	_, err = newOrm.Delete(&article)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/showArticle", 302)

}

// @router /updateById [post]
func (c *ArticleController) UpdateById() {
	id, err := c.GetInt("id")
	if err != nil {
		logs.Error(err)
	}
	title := c.GetString("title")
	typeName := c.GetString("type")
	content := c.GetString("content")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(title, "title").Message("title不能为空")
	valid.Required(typeName, "type").Message("type不能为空")
	valid.Required(content, "content").Message("content不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
		}
		c.TplName = "update.html"
		return
	}

	o := orm.NewOrm()
	articleType := models.ArticleType{}
	err = o.QueryTable("ArticleType").Filter("TypeName", typeName).One(&articleType)
	if err != nil {
		logs.Error(err)
	}
	article := models.Article{Id: id, Title: title, Type: &articleType, Content: content}
	file, header, err := c.GetFile("img")
	cols := []string{"title", "type", "content"}
	if err != http.ErrMissingFile {
		if err != nil {
			logs.Error(err)
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				logs.Error(err)
			}
		}(file)

		ext := path.Ext(header.Filename)
		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			logs.Error("上传文件格式错误")
			return
		}
		if header.Size > 500000000000 {
			logs.Error("文件过大")
			return
		}
		timeFormat := time.Now().Format("2006-01-02 15-04-05")
		imgPath := "static/img/" + timeFormat + ext
		err = c.SaveToFile("img", imgPath)
		if err != nil {
			logs.Error(err)
			return
		}
		article.Img = imgPath
		cols = append(cols, "img")
	}

	_, err = o.Update(&article, cols...)
	if err != nil {
		logs.Error(err)
		return
	}

	c.Redirect("/showArticle", 302)

}

func NextPage(page int) int {
	return page + 1
}

func PrePage(page int) int {
	return page - 1
}
