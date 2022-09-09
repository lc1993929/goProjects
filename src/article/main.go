package main

import (
	_ "article/models"
	_ "article/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
