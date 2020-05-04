package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/google/uuid"
	"github.com/lescactus/http-gallery-beego/controllers"
)

var requestID = func(ctx *context.Context) {

}

func init() {
	beego.SetStaticPath("static/css", "static/css")
	beego.SetStaticPath("static/img", "static/img")
	beego.SetStaticPath("static/js", "static/js")
	beego.SetStaticPath("static/fonts", "static/fonts")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})

	beego.InsertFilter("*", beego.BeforeExec, func(ctx *context.Context) {
		uuid := uuid.New()
		ctx.Input.SetData("requestid", "req id: "+uuid.String()+" - ")
	})

}
