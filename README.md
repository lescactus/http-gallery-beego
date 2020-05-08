HTTP Gallery
==================

A Beego web app that allow users to upload local images that are displayed in a html gallery and carousel. This is the rewriting in Go of one of my previous mini web project: **[http_gallery](https://github.com/lescactus/http_gallery)**.

**gallery**: [gallery-beego.alexasr.tk][1]

It uses:

   * [Go][2]
   * [Beego][3]
   * [Bootstrap v3][4]
   * [JQuery v2.2.4][5]
   * [ludovicscribe/bootstrap-gallery][6]
   * [kartik-v/bootstrap-fileinput][7]

Use it now
----------

```sh
# Install http-gallery-beego
$ go get github.com/lescactus/http-gallery-beego

# Run the webserver
$ http-gallery-beego
2020/05/08 15:59:10.452 [I]  Directory uploads/ is not present. Creating it...
2020/05/08 15:59:10.452 [I]  Directory thumbnails/ is not present. Creating it...
2020/05/08 15:59:10.452 [I]  No HTTP_PORT environment variable provided. Fallback to :8080
2020/05/08 15:59:10.452 [I]  No XSRF_KEY environment variable provided. A default one will be randomly generated
2020/05/08 15:59:10.452 [I]  No XRSF_EXPIRE environment variable privided. Fallback to 0
2020/05/08 15:59:10.467 [I]  http server Running on http://:8080

# Now point your browser at http://127.0.0.1:8080
```
#### Configuration
You can use the following environment variables

* `HTTP_PORT`: The tcp port the web server will listen to. Must be an integer (default: `8080`).
* `XSRF_KEY`: The XSRF key used by Beego (https://beego.me/docs/mvc/controller/xsrf.md). If not provided, a default one will be randomly generated.
* `XRSF_EXPIRE`: The XSRF expiration time. Must be an integer (default `).
* `XSRF_KEY_PATH`: If set, Beego will look into the file located at `XSRF_KEY_PATH` to read the XSRF key from (Example: `/secret/xsrf_key`). Useful if mounted from a Kubernetes secret.

### Docker
**gallery** can easily be dockerized and is shipped with a ``Dockerfile``.

By default, the container will expose port 8080, so change this within the ``Dockerfile`` if necessary. When ready, simply use the ``Dockerfile`` to build the image.

```sh
cd app
docker build -t gallery .
```
This will create the Docker image.

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 80 of the host to port 8080 of the container:

```sh
docker run -d -p 80:8080 --restart="always" --name gallery gallery 
```

Now point your browser at http://127.0.0.1/ 

### Docker volumes
Since Docker is stateless, uploaded files are removed when the container is destroyed. You can make your data persistent by mounting the `uploads/` and `thumbnails/` folders in Docker volumes:
```sh
# docker run
docker run -d \
-p 80:8080 \
-v $(pwd)/uploads:/app/uploads \
-v $(pwd)/thumbnails:/app/thumbnails \
--name app \
gallery
```




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



[1]: http://gallery.alexasr.tk/
[2]: https://golang.org/
[3]: https://beego.me/
[4]: https://getbootstrap.com/
[5]: https://jquery.com/
[6]: https://github.com/ludovicscribe/bootstrap-gallery
[7]: https://github.com/kartik-v/bootstrap-fileinput