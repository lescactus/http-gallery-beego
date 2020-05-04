package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lescactus/http-gallery-beego/routers"
)

func main() {
	beego.Run()
}
