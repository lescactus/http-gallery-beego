package routers

import (
	"os"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/google/uuid"
	"github.com/lescactus/http-gallery-beego/controllers"
	"github.com/lescactus/http-gallery-beego/models"
)

func createDirectoryIfNotPresent(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		logs.Info("Directory " + dirPath + " is not present. Creating it...")
		if err = os.MkdirAll(dirPath, os.FileMode(0750)); err != nil {
			logs.Critical("Error: " + err.Error())
			os.Exit(1)
		}
	}
}

func init() {
	// If local filesystem is the chosen storage backend, create the upload and thumbnails directories if needed
	if models.StorageType == "local" {
		createDirectoryIfNotPresent(models.UploadDirectory)
		createDirectoryIfNotPresent(models.ThumbnailsDirectory)
	}

	// If GCP bucket is the chosen storage backend, create a temporary directory if needed
	if models.StorageType == "GCP" {
		createDirectoryIfNotPresent(models.TmpDirectory)
	}

	// If local filesystem is the chosen storage backend, set the uploads and thumbnails directory as a static path
	if models.StorageType == "local" {
		beego.SetStaticPath(models.UploadDirectory, models.UploadDirectory)
		beego.SetStaticPath(models.ThumbnailsDirectory, models.ThumbnailsDirectory)
	}
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
