package controllers

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/api/iterator"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/lescactus/http-gallery-beego/models"

	"cloud.google.com/go/storage"
)

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

// Return true if file is a real image and false if not
func isAnImage(f fs.DirEntry) bool {
	if !f.IsDir() {
		f, err := os.Open(f.Name())
		if err != nil {
			return false
		}

		contentType, err := getFileContentType(f)
		if err != nil {
			return false
		}

		return isContentTypeAllowed(contentType)
	}
	return false
}

// Generate the thumbnail name from the original image
// by appending '-thumb' before the extension name
// Ex: test.jpg => test-thumb.jpg
func generateThumbnailName(orig string) string {
	return strings.Trim(orig, filepath.Ext(orig)) + "-thumb" + filepath.Ext(orig)
}

func errorHandler(msg string, c *MainController, err error, flash *beego.FlashData) {
	logs.Error(c.Ctx.Input.GetData("requestid"), err.Error())
	flash.Error(msg)
	flash.Store(&c.Controller)
	c.Redirect("/", 302)
}

// Upload source into a Google Storage bucket
func uploadGoogleStorage(source, destination string) error {
	ctx := context.Background()

	file, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("Failed to open file: " + err.Error())
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("Failed to create a Google Cloud Storage NewClient(): " + err.Error())
	}

	// Write object in the bucket
	writer := client.Bucket(models.BucketName).Object(destination).NewWriter(ctx)
	if _, err = io.Copy(writer, file); err != nil {
		return fmt.Errorf("Failed to write to bucket: " + err.Error())
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("Failed to close writer: " + err.Error())
	}

	return nil
}

func getBucketFiles() (map[string]string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	delim := "/"
	listUploads := client.Bucket(models.BucketName).Objects(ctx, &storage.Query{
		Prefix:    models.UploadDirectory,
		Delimiter: delim,
	})
	listThumbnails := client.Bucket(models.BucketName).Objects(ctx, &storage.Query{
		Prefix:    models.ThumbnailsDirectory,
		Delimiter: delim,
	})
	images := map[string]string{}
	for {
		attrsUpload, err := listUploads.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		attrsThumbnail, err := listThumbnails.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		images[filepath.Base(attrsUpload.Name)] = filepath.Base(attrsThumbnail.Name)
	}

	return images, err
}
