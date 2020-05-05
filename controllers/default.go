package controllers

import (
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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
	uploadDirectory = "uploads/"

	// THumbnails directiry
	thumbnailDirectory = "thumbnails/"

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

// Return true if file is a real file and false if it is a directory
func isAFile(file string) bool {
	i, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}

	return !i.IsDir()
}

// Generate the thumbnail name from the original image
// by appending '-thumb' before the extension name
// Ex: test.jpg => test-thumb.jpg
func generateThumbnailName(orig string) string {
	return strings.Trim(orig, filepath.Ext(orig)) + "-thumb" + filepath.Ext(orig)
}

// Get handle GET requests
func (c *MainController) Get() {
	beego.ReadFromRequest(&c.Controller)

	// Get all saved images and thumbnails
	images := map[string]string{}
	files, err := ioutil.ReadDir(uploadDirectory)
	if err != nil {
		logs.Critical("Error: " + err.Error())
	}
	for _, file := range files {
		if isAFile(uploadDirectory + file.Name()) {
			images[file.Name()] = generateThumbnailName(file.Name())
		}
	}

	c.Data["uploadDirectory"] = uploadDirectory
	c.Data["thumbnailDirectory"] = thumbnailDirectory
	c.Data["images"] = images
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["htmlInputName"] = htmlInputName
	c.TplName = "upload.tpl"
}

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
	if err := c.SaveToFile("file", uploadDirectory+fileName); err != nil {
		errorHandler("Error while saving "+uploadDirectory+fileName+" on the local filesystem", c, err, flash)
		return

	}

	// Create and save the thumbnail from the uploaded image
	thumbnailName := generateThumbnailName(fileName)
	thumbnail, err := imaging.Open(uploadDirectory + fileName)
	if err != nil {
		errorHandler("Error while creating thumbnail for "+fileName, c, err, flash)
		return
	}

	thumbnail = imaging.Fill(thumbnail, 300, 300, imaging.Center, imaging.Box)
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
