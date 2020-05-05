package controllers

import (
	"html/template"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/disintegration/imaging"
)

var (
	// Content type allowed to be save and stored. Only images
	allowedContentTypes = []string{"image/jpeg", "image/jpg", "application/jpeg", "application/jpg", "image/png", "image/ief"}

	// Upload directory
	uploadDirectory = "./uploads/"

	// THumbnails directiry
	thumbnailDirectory = "./thumbnails/"

	// Name of the type="file" html input
	htmlInputName = "file"
)

type MainController struct {
	beego.Controller
}

// GetFileContentType return the Content Type of a file sent through a POST request
// Inspired by https://golangcode.com/get-the-content-type-of-file/
func getFileContentType(file multipart.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// Return true if the content-type in argument match one in allowedMimeTypes
func isContentTypeAllowed(contentType string) bool {
	for _, value := range allowedContentTypes {
		if contentType == value {
			return true
		}
	}
	return false
}

func (c *MainController) Get() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["htmlInputName"] = htmlInputName
	c.TplName = "upload.tpl"
}

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
	if err := c.SaveToFile("file", uploadDirectory+fileName); err != nil {
		errorHandler("Error while saving "+uploadDirectory+fileName+" on the local filesystem", c, err, flash)
		return

	}

	// Create and save the thumbnail from the uploaded image
	thumbnailName := strings.Trim(fileName, filepath.Ext(fileName)) + "-thumb" + filepath.Ext(fileName)
	logs.Debug("Filename: " + fileName + ", Thumbnail: " + thumbnailName)
	thumbnail, err := imaging.Open(uploadDirectory + fileName)
	if err != nil {
		errorHandler("Error while creating thumbnail for "+fileName, c, err, flash)
		return
	}

	thumbnail = imaging.Resize(thumbnail, 300, 300, imaging.Lanczos)
	err = imaging.Save(thumbnail, thumbnailDirectory+thumbnailName)
	if err != nil {
		errorHandler("Error while creating thumbnail for "+fileName, c, err, flash)
		return
	}

	logs.Info(c.Ctx.Input.GetData("requestid"), "New file saved successfully: "+thumbnailName)
	flash.Success("File successfully uploaded")

	flash.Store(&c.Controller)
	c.Redirect("/", 301)
	return
	//c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	//c.TplName = "upload.tpl"
}

func errorHandler(msg string, c *MainController, err error, flash *beego.FlashData) {
	logs.Error(c.Ctx.Input.GetData("requestid"), err.Error())
	flash.Error(msg)
	flash.Store(&c.Controller)
	c.Redirect("/", 302)
	return
}
