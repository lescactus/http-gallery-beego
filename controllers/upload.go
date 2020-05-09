package controllers

import (
	"github.com/astaxie/beego"
)

var (
	// Content type allowed to be save and stored. Only images
	allowedContentTypes = []string{"image/jpeg", "image/jpg", "application/jpeg", "application/jpg", "image/png", "image/ief"}

	// Name of the type="file" html input
	htmlInputName = "file"

	// Name of the theme cookie
	themeCookie = "theme"

	// Name of the default theme
	defaultTheme = "flatty"
)

type MainController struct {
	beego.Controller
}
