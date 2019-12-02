package main

import (
	"github.com/astaxie/beego"
	_ "github.com/guxiaogang/SecKill-Proxy/router"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = false
	err := initConfig()
	if err != nil {
		panic(err)
		return
	}
	err = initSec()
	if err != nil {
		panic(err)
		return
	}
	beego.Run()
}
