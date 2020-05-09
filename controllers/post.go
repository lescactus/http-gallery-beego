package controllers

import (
	"os"

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

	imagePath := ""
	thumbnailName := ""
	thumbnailPath := ""

	if models.StorageType == "local" {
		imagePath = models.UploadDirectory + fileName
		thumbnailName = generateThumbnailName(fileName)
		thumbnailPath = models.ThumbnailsDirectory + thumbnailName
	} else { // GCP bucket
		imagePath = models.TmpDirectory + fileName
		thumbnailName = generateThumbnailName(fileName)
		thumbnailPath = models.TmpDirectory + thumbnailName
	}

	// Save file in the local filesystem and create a new thumbnail out of it
	if err := c.SaveToFile("file", imagePath); err != nil {
		errorHandler("Error while saving "+imagePath+" on the local filesystem", c, err, flash)
		return
	}
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file saved successfully: "+fileName)

	// Create and save the thumbnail from the uploaded image
	thumbnail, err := imaging.Open(imagePath)
	if err != nil {
		errorHandler("Error while creating thumbnail for "+fileName, c, err, flash)
		return
	}

	thumbnail = imaging.Fill(thumbnail, 300, 300, imaging.Center, imaging.Box)
	err = imaging.Save(thumbnail, thumbnailPath)
	if err != nil {
		errorHandler("Error while creating thumbnail for "+fileName, c, err, flash)
		return
	}
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file saved successfully: "+thumbnailName)

	if models.StorageType == "local" {
		flash.Success("File successfully uploaded")

		flash.Store(&c.Controller)
		c.Redirect("/", 301)
		return
	}

	// Push image and thumbnail to Google Storage
	if models.StorageType == "GCP" {
		defer os.Remove(imagePath)
		defer os.Remove(thumbnailPath)

		if err := uploadGoogleStorage(imagePath, models.UploadDirectory+fileName); err != nil {
			errorHandler("Error while saving "+fileName+" to Google Storage Bucket", c, err, flash)
		}
		if err := uploadGoogleStorage(thumbnailPath, models.ThumbnailsDirectory+thumbnailName); err != nil {
			errorHandler("Error while saving "+thumbnailName+" to Google Storage Bucket", c, err, flash)
		}

		flash.Success("File successfully uploaded")

		flash.Store(&c.Controller)
		c.Redirect("/", 301)
		return
	}
}
