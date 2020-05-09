package models

var (
	// UploadDirectory is the name of the directory to store the images into
	UploadDirectory = "uploads/"

	// ThumbnailsDirectory is the name of the directory to store the thumbnails into
	ThumbnailsDirectory = "thumbnails/"

	// StorageType contains the type of storage backend. Either local or GCP
	StorageType = ""

	// BucketName is the name of the GCP bucket
	BucketName = ""
)
