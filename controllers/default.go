package controllers

import (
	"html/template"
	"mime/multipart"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	// Content type allowed to be save and stored. Only images
	allowedContentTypes = []string{"image/jpeg", "image/jpg", "application/jpeg", "application/jpg", "image/png", "image/ief"}

	// Upload directory
	uploadDirectory = "./uploads/"

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
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["htmlInputName"] = htmlInputName
	c.TplName = "upload.tpl"
}

func (c *MainController) Post() {
	// Get
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
		logs.Error(c.Ctx.Input.GetData("requestid"), "Can't get "+fileName+"Content-Type")
	}
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file Content-Type: "+contentType)

	// The file sent isn't allowed to be uploaded
	if !isContentTypeAllowed(contentType) {
		logs.Error(c.Ctx.Input.GetData("requestid"), "Content-Type of "+fileName+" isn't allowed")
		c.Ctx.Output.Body([]byte("Content-Type of " + fileName + " isn't allowed"))
		return
	}
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file Content-Type belongs to the allowed Content-Type")

	// Save file in the local filesystem
	if err := c.SaveToFile("file", uploadDirectory+fileName); err != nil {
		logs.Error(c.Ctx.Input.GetData("requestid"), err.Error())
		return

	}
	logs.Info(c.Ctx.Input.GetData("requestid"), "New file saved successfully: "+fileName)

	c.Redirect("/", 301)
	return
	//c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	//c.TplName = "upload.tpl"
}
