package controllers

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
func isAnImage(file string) bool {
	i, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}

	if !i.IsDir() {
		f, err := os.Open(file)
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
