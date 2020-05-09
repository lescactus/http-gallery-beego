package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/lescactus/http-gallery-beego/models"
	_ "github.com/lescactus/http-gallery-beego/routers"
)

var (
	// HTTPPortEnvVariable is the name of the HttpPort environment variable
	HTTPPortEnvVariable = "HTTP_PORT"

	// XSRFKeyEnvVariable is the name of the XSRFKey environment variable
	XSRFKeyEnvVariable = "XSRF_KEY"

	// XSRFExpireEnvVariable is the name of the XSRFExpire environment variable
	XSRFExpireEnvVariable = "XRSF_EXPIRE"

	// XSRFKeyPathEnvVariable is the path to a file containig the XSRFKey
	XSRFKeyPathEnvVariable = "XSRF_KEY_PATH"

	// StorageTypeEnvVariable is the name of the environment variable containing the storage type (local or GCP)
	StorageTypeEnvVariable = "STORAGE_TYPE"

	// GCPBucketNameEnvVariable is the name of the environment variable containing the name of the GCP bucket to store images to
	GCPBucketNameEnvVariable = "BUCKET_NAME"
)

func generteRandomXSRFKey() string {
	charset := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 64

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	// Define HTTPPort
	if os.Getenv(HTTPPortEnvVariable) != "" {
		var err error
		beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(os.Getenv(HTTPPortEnvVariable))
		if err != nil {
			logs.Critical("Error: " + HTTPPortEnvVariable + " must be an integer ! " + err.Error())
			os.Exit(1)
		}
	} else {
		logs.Info("No " + HTTPPortEnvVariable + " environment variable provided. Fallback to :8080")
		beego.BConfig.Listen.HTTPPort = 8080
	}

	// Check if XSRFKeyPathEnvVariable is set, if yes, read the XSRFKey from this file, otherwise,
	// read it from XSRFKeyEnvVariable if set, otherwise generate it randomly
	if os.Getenv(XSRFKeyPathEnvVariable) != "" {
		XSRFKey, err := ioutil.ReadFile(os.Getenv(XSRFKeyPathEnvVariable))
		if err != nil {
			logs.Critical("Error: " + XSRFKeyPathEnvVariable + " can't be read: " + err.Error())
			os.Exit(1)
		}
		beego.BConfig.WebConfig.XSRFKey = string(XSRFKey)
	} else if os.Getenv(XSRFKeyEnvVariable) != "" {
		beego.BConfig.WebConfig.XSRFKey = os.Getenv(XSRFKeyEnvVariable)
	} else {
		logs.Info("No " + XSRFKeyEnvVariable + " environment variable provided. A default one will be randomly generated")
		beego.BConfig.WebConfig.XSRFKey = generteRandomXSRFKey()
	}

	// Define XSRFExpire
	if os.Getenv(XSRFExpireEnvVariable) != "" {
		var err error
		beego.BConfig.WebConfig.XSRFExpire, err = strconv.Atoi(os.Getenv(XSRFExpireEnvVariable))
		if err != nil {
			logs.Critical("Error: " + XSRFExpireEnvVariable + " must be an integer ! " + err.Error())
			os.Exit(1)
		}

	} else {
		logs.Info("No " + XSRFExpireEnvVariable + " environment variable privided. Fallback to 0")
		beego.BConfig.WebConfig.XSRFExpire = 0
	}

	// Define if storage backend is local or GCP bucket. In case it's GCP, get the bucket name
	if os.Getenv(StorageTypeEnvVariable) == "" {
		logs.Info("No " + StorageTypeEnvVariable + " environment variable provided. Fallback to 'local'")
		models.StorageType = "local"
	} else if os.Getenv(StorageTypeEnvVariable) == "local" {
		models.StorageType = "local"
	} else if os.Getenv(StorageTypeEnvVariable) == "GCP" {
		if os.Getenv(GCPBucketNameEnvVariable) != "" {
			models.StorageType = "GCP"
			models.BucketName = os.Getenv(GCPBucketNameEnvVariable)
		} else {
			logs.Error("When " + StorageTypeEnvVariable + " is set to GCP, " + GCPBucketNameEnvVariable + " must not be empty.")
			os.Exit(1)
		}

	} else {
		logs.Error(StorageTypeEnvVariable + " must either be 'local' or 'GCP'. Got " + os.Getenv(StorageTypeEnvVariable) + ". Fallback to 'local'")
		models.StorageType = "local"
	}

	beego.Run()
}
