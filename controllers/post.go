package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/disintegration/imaging"
	"github.com/lescactus/http-gallery-beego/models"
)

// Post handle POST requests (form submittion)
// Verify that the file submitted is a real image
// Save and create a thumbnail out of it
func (c *MainController) Post() {
	flash := beego.NewFlash()

	// Get file from HTML form
	file, header, err := c.GetFile("file")
	if err != nil {
		logs.Error(err.Error())
		return
	}

	fileName := header.Filename
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file received: "+fileName)

	// Get the Content-Type of the file sent using the html form
	contentType, err := getFileContentType(file)
	if err != nil {
		errorHandler("Error while getting content type of "+fileName, c, err, flash)
		return
	}
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file Content-Type: "+contentType)

	// The file sent isn't allowed to be uploaded
	if !isContentTypeAllowed(contentType) {
		errorHandler("Error: File type of "+fileName+" is not allowed", c, err, flash)
		return
	}
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file Content-Type belongs to the allowed Content-Type")

	// Save file in the local filesystem and create a new thumbnail out of it
	if err := c.SaveToFile("file", models.UploadDirectory+fileName); err != nil {
		errorHandler("Error while saving "+models.UploadDirectory+fileName+" on the local filesystem", c, err, flash)
		return

	}

	// Create and save the thumbnail from the uploaded image
	thumbnailName := generateThumbnailName(fileName)
	thumbnail, err := imaging.Open(models.UploadDirectory + fileName)
	if err != nil {
		errorHandler("Error while creating thumbnail for "+fileName, c, err, flash)
		return
	}

	thumbnail = imaging.Fill(thumbnail, 300, 300, imaging.Center, imaging.Box)
	err = imaging.Save(thumbnail, models.ThumbnailsDirectory+thumbnailName)
	if err != nil {
		errorHandler("Error while creating thumbnail for "+fileName, c, err, flash)
		return
	}

	logs.Info(c.Ctx.Input.GetData("requestid"), "New file saved successfully: "+thumbnailName)
	flash.Success("File successfully uploaded")

	flash.Store(&c.Controller)
	c.Redirect("/", 301)
	return
}

func errorHandler(msg string, c *MainController, err error, flash *beego.FlashData) {
	logs.Error(c.Ctx.Input.GetData("requestid"), err.Error())
	flash.Error(msg)
	flash.Store(&c.Controller)
	c.Redirect("/", 302)
	return
}
