package controllers

import (
	"html/template"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/lescactus/http-gallery-beego/models"
)

// Get handle GET requests
func (c *MainController) Get() {
	beego.ReadFromRequest(&c.Controller)

	// Get name of the current theme in a cookie.
	// If empty, set the default theme
	theme := ""
	if c.Ctx.GetCookie(themeCookie) != "" {
		theme = c.Ctx.GetCookie(themeCookie)
	} else {
		theme = defaultTheme
	}

	// Get all saved images and thumbnails
	images := map[string]string{}
	if models.StorageType == "local" {
		files, err := ioutil.ReadDir(models.UploadDirectory)
		if err != nil {
			logs.Critical("Error: " + err.Error())
		}
		for _, file := range files {
			if isAnImage(models.UploadDirectory + file.Name()) {
				images[file.Name()] = generateThumbnailName(file.Name())
			}
		}
	} else {
		var err error
		images, err = getBucketFiles()
		if err != nil {
			logs.Critical("Error: " + err.Error())
			return
		}
	}

	if models.StorageType == "local" {
		c.Data["uploadDirectory"] = models.UploadDirectory
		c.Data["thumbnailDirectory"] = models.ThumbnailsDirectory
	} else {
		c.Data["uploadDirectory"] = "https://storage.googleapis.com/" + models.BucketName + "/" + models.UploadDirectory
		c.Data["thumbnailDirectory"] = "https://storage.googleapis.com/" + models.BucketName + "/" + models.ThumbnailsDirectory
	}
	c.Data["images"] = images
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["htmlInputName"] = htmlInputName
	c.Data["theme"] = theme
	c.TplName = "upload.tpl"
}
