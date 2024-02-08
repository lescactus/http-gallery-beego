# HTTP Gallery

[![cloud_build_status](https://storage.googleapis.com/gcp-build-badge/http-gallery-beego/statusbadge.svg)](https://github.com/lescactus/http-gallery-beego) [![Go build and test](https://github.com/lescactus/http-gallery-beego/actions/workflows/go.yaml/badge.svg)](https://github.com/lescactus/http-gallery-beego/actions/workflows/go.yaml)

![go](https://img.shields.io/badge/go->=1.16-blue)

A Beego web app that allow users to upload local images that are displayed in a html gallery and carousel. This is the rewriting in Go of one of my previous mini web project: **[http_gallery](https://github.com/lescactus/http_gallery)**.

**gallery**: [gallery-beego.alexasr.fr][1]

It uses:

   * [Go][2]
   * [Beego][3]
   * [Bootstrap v3][4]
   * [JQuery v2.2.4][5]
   * [ludovicscribe/bootstrap-gallery][6]
   * [kartik-v/bootstrap-fileinput][7]

## Requirements

* Golang 1.16 or higher

## Getting started

Prebuilt binaries can be downloaded from the GitHub Releases [section](https://github.com/lescactus/http-gallery-beego/releases), or using a Docker image from the Github Container Registry. See [here](https://github.com/lescactus/http-gallery-beego#running-with-docker)

### With `go get`

```sh
# Install http-gallery-beego
$ go get github.com/lescactus/http-gallery-beego

# Run the webserver
$ ~/go/bin/http-gallery-beego
2020/05/09 15:17:31.081 [I]  No STORAGE_TYPE environment variable provided. Fallback to 'local'
2020/05/09 15:17:31.081 [I]  Directory uploads/ is not present. Creating it...
2020/05/09 15:17:31.081 [I]  Directory thumbnails/ is not present. Creating it...
2020/05/09 15:17:31.081 [I]  No HTTP_PORT environment variable provided. Fallback to :8080
2020/05/09 15:17:31.081 [I]  No XSRF_KEY environment variable provided. A default one will be randomly generated
2020/05/09 15:17:31.081 [I]  No XRSF_EXPIRE environment variable privided. Fallback to 0
2020/05/09 15:17:31.085 [I]  http server Running on http://:8080

# Now point your browser at http://127.0.0.1:8080
```

### With docker

**gallery** can easily be dockerized and is shipped with a `Dockerfile`.

By default, the container will expose port 8080. When ready, simply use the `Dockerfile` to build the image.

```sh
docker build -t gallery .
```
This will create the Docker image.

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 80 of the host to port 8080 of the container:

```sh
docker run -d -p 80:8080 --restart="always" --name gallery gallery 
```

Now point your browser at http://127.0.0.1/

## Building

<details>

### From source with Go

You need a working [go](https://golang.org/doc/install) toolchain (It has been developped and tested with go 1.16 and should work with go >= 1.16). Refer to the official documentation for more information (or from your Linux/Mac/Windows distribution documentation to install it from your favorite package manager).

```bash
# Clone this repository
git clone https://github.com/lescactus/http-gallery-beego.git && cd http-gallery-beego/

# Build from sources. Use the '-o' flag to change the compiled binary name
go build

# Default compiled binary is http-gallery-beego
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./http-gallery-beego
```

### From source with docker

If you don't have [go](https://golang.org/doc/install) installed but have docker, run the following command to build inside a docker container:

```bash
# Build from sources inside a docker container. Use the '-o' flag to change the compiled binary name
# Warning: the compiled binary belongs to root:root
docker run --rm -it -v "$PWD":/app -w /app golang:1.16 go build

# Default compiled binary is http-gallery-beego
# You can optionnaly move it somewhere in your $PATH to access it shell wide
./http-gallery-beego
```

The server is accessible at http://127.0.0.1:8080

### With Docker

`http-gallery-beego` comes with a `Dockerfile`. To build the image:

```bash
docker build -t http-gallery-beego .

docker run -d -p 8080:8080 --restart="always" --name http-gallery-beego http-gallery-beego 
```

The server is accessible at http://127.0.0.1:8080

</details>

## Configuration

You can use the following environment variables

* `HTTP_PORT`: The tcp port the web server will listen to. Must be an integer (default: `8080`).
* `XSRF_KEY`: The XSRF key used by Beego (https://beego.me/docs/mvc/controller/xsrf.md). If not provided, a default one will be randomly generated.
* `XRSF_EXPIRE`: The XSRF expiration time. Must be an integer (default `).
* `XSRF_KEY_PATH`: If set, Beego will look into the file located at `XSRF_KEY_PATH` to read the XSRF key from (Example: `/secret/xsrf_key`). Useful if mounted from a Kubernetes secret.
* `STORAGE_TYPE`: Can be either `local` or `GCP` (default: `local`). If set to `local`, the images and thumbnails will be stored on the local filesystem inside `./uploads/` and `./thumbnails/` directories (they will be created if non-exsitent during start-up). If set to `GCP`, the images and thumbnails will be stored in the `BUCKET_NAME` [Google Cloud Storage Bucket](https://cloud.google.com/storage/docs/json_api/v1/buckets).
* `BUCKET_NAME`: Name of the [Google Cloud Storage Bucket](https://cloud.google.com/storage/docs/json_api/v1/buckets) used to store images and thumbnails. It has no effect if `STORAGE_TYPE` is set to `local`. 

## Persistence

### `local` storage type 

When `STORAGE_TYPE` is set to `local`, the images and thumbnails will be written on the local filesystem inside `./uploads/` and `./thumbnails/` directories. If these directories do not exist during the application start-up, they will be created. You must ensure that the user running the application has the proper write permissions.

### Docker volumes

Since Docker is stateless, uploaded files are lost when the container is destroyed. You can make your data persistent by mounting the `uploads/` and `thumbnails/` folders in Docker volumes:

```sh
# docker run
docker run \
   --detach \
   --publish 80:8080 \
   --env STORAGE_TYPE="local" \
   --volume $(pwd)/uploads:/app/uploads \
   --volume $(pwd)/thumbnails:/app/thumbnails \
   --name gallery \
   gallery
```

### `GCP` storage type

When `STORAGE_TYPE` is set to `GCP`, the images and thumbnails will be written in the `BUCKET_NAME` [Google Cloud Storage Bucket](https://cloud.google.com/storage/docs/json_api/v1/buckets).
The application will not create the bucket and will throw an error if it don't exist or of the proper Google Cloud credentials given are not enough.
You must ensure the file inside this bucket are publicly accessible from the Internet. [Ex](https://cloud.google.com/storage/docs/access-control/making-data-public#buckets):

* Create a GCP bucket with Uniform bucket-level access
* In the permissions tab of the created bucket, click the **Add members** button
* In the **New members** field, add `allUsers`
* Select the role **Storage Object Viewer** from the **Cloud Storage** sub-menu

#### Required permissions
Once the public bucket is created, you must ensure to set-up the proper write permissions on this bucket. You can create a new service account or an existing one and use its credentials to authenticate to the Google Cloud API (https://cloud.google.com/docs/authentication/production#obtaining_and_providing_service_account_credentials_manually). 
* [Create a new service account](https://cloud.google.com/docs/authentication/getting-started#creating_a_service_account)
* In the **Role** list, chose `Storage Admin` from the **Cloud Storage** sub-menu
* Create and save the service account JSON key locally
* Export the `GOOGLE_APPLICATION_CREDENTIALS` environment variable with the path of the saved key (ex: `export GOOGLE_APPLICATION_CREDENTIALS=$HOME/.gcloud/serviceaccount.json`)

#### Docker
To use the service account key from inside a docker container, you must bind mount the key within a docker volume:
```sh
docker run \
   --detach \
   --publish 80:8080 \
   --name gallery \
   --env STORAGE_TYPE="GCP" \
   --env BUCKET_NAME="my-bucket-name" \
   --env GOOGLE_APPLICATION_CREDENTIALS="/tmp/creds" \
   --volume "$HOME/.gcloud/test.json:/tmp/creds" gallery
```

### Google Cloud Run

If you plan to run this application with [Cloud Run](https://cloud.google.com/run), you must ensure your Cloud Run service account can also write to your Cloud Storage Bucket. 
Add the role **Storage Object Creator** to your service account (default service account name is `xxxxxxxxxx-compute@developer.gserviceaccount.com`)
Simply add the `STORAGE_TYPE=GCP` and `BUCKET_NAME=my-bucket-name` environment variables.

Screenshots
-----------
**Index**
![Index](https://i.imgur.com/DIMzgU6.png "Index")
***
**Image upload form**
![Upload an image](https://i.imgur.com/RGCiG8l.png "Upload an image")
***
**Gallery**
![Gallery](https://i.imgur.com/eadFN3J.png "Gallery")
***
**Carousel**
![Carousel](https://i.imgur.com/WaMuiv9.png "Carousel")
***
**Responsive gallery**
![Responsive](https://i.imgur.com/fGxH2CH.png "Responsive")
***



[1]: https://gallery-beego.alexasr.fr/
[2]: https://golang.org/
[3]: https://beego.me/
[4]: https://getbootstrap.com/
[5]: https://jquery.com/
[6]: https://github.com/ludovicscribe/bootstrap-gallery
[7]: https://github.com/kartik-v/bootstrap-fileinput
