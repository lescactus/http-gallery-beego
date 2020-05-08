FROM library/golang as builder

# Recompile the standard library without CGO
RUN go get -v "github.com/astaxie/beego" "github.com/google/uuid" "github.com/disintegration/imaging" \
  && CGO_ENABLED=0 go install -v -a std

ENV APP_DIR $GOPATH/src/github.com/lescactus/http-gallery-beego
RUN mkdir -p $APP_DIR

ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s' -o main

FROM alpine

WORKDIR /app

RUN chown -R 65534:65534 /app

COPY --from=builder --chown=65534:65534 /go/src/github.com/lescactus/http-gallery-beego/main /app
COPY --from=builder --chown=65534:65534 /go/src/github.com/lescactus/http-gallery-beego/views /app/views
COPY --from=builder --chown=65534:65534 /go/src/github.com/lescactus/http-gallery-beego/static /app/static

EXPOSE 8080

# nobody
USER 65534

CMD ["./main"]
