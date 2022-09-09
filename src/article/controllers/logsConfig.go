package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

func init() {
	err := logs.SetLogger(logs.AdapterConsole)
	if err != nil {
		fmt.Println(err)
	}
	logs.Async()

}
