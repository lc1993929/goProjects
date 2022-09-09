package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int
	UserName string
	Password string
	Articles []*Article `orm:"rel(m2m)"`
}

type Article struct {
	Id      int          `orm:"auto"`
	Title   string       `valid:"Required"`
	Img     string       `valid:"Required"`
	Content string       `valid:"Required"`
	Time    time.Time    `orm:"type(datetime);auto_now_add"`
	Count   int          `orm:"default(0)"`
	Type    *ArticleType `valid:"Required" orm:"rel(fk)"`
	Users   []*User      `orm:"reverse(many)"`
}

type ArticleType struct {
	Id       int
	TypeName string     `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error(err)
	}
}
