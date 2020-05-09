package models

import (
	"os"

	"github.com/astaxie/beego/logs"
)

var (
	// StorageTypeEnvVariable is the name of the environment variable containing the storage type (local or GCP)
	StorageTypeEnvVariable = "STORAGE_TYPE"

	// GCPBucketNameEnvVariable is the name of the environment variable containing the name of the GCP bucket to store images to
	GCPBucketNameEnvVariable = "BUCKET_NAME"

	// UploadDirectory is the name of the directory to store the images into
	UploadDirectory = "uploads/"

	// ThumbnailsDirectory is the name of the directory to store the thumbnails into
	ThumbnailsDirectory = "thumbnails/"

	// TmpDirectory is the name of the temporary directory
	TmpDirectory = "tmp/"

	// StorageType contains the type of storage backend. Either local or GCP
	StorageType string

	// BucketName is the name of the GCP bucket
	BucketName string
)

func init() {
	// Define if storage backend is local or GCP bucket. In case it's GCP, get the bucket name
	if os.Getenv(StorageTypeEnvVariable) == "" {
		logs.Info("No " + StorageTypeEnvVariable + " environment variable provided. Fallback to 'local'")
		StorageType = "local"
	} else if os.Getenv(StorageTypeEnvVariable) == "local" {
		StorageType = "local"
	} else if os.Getenv(StorageTypeEnvVariable) == "GCP" {
		if os.Getenv(GCPBucketNameEnvVariable) != "" {
			StorageType = "GCP"
			BucketName = os.Getenv(GCPBucketNameEnvVariable)
		} else {
			logs.Error("When " + StorageTypeEnvVariable + " is set to GCP, " + GCPBucketNameEnvVariable + " must not be empty.")
			os.Exit(1)
		}

	} else {
		logs.Error(StorageTypeEnvVariable + " must either be 'local' or 'GCP'. Got " + os.Getenv(StorageTypeEnvVariable) + ". Fallback to 'local'")
		StorageType = "local"
	}
}
