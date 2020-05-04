FROM library/golang as builder

RUN go get "github.com/astaxie/beego" "github.com/google/uuid"

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/lescactus/http-gallery-beego
RUN mkdir -p $APP_DIR

ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s' -o main

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/github.com/lescactus/http-gallery-beego/main /app
COPY --from=builder /go/src/github.com/lescactus/http-gallery-beego/views /app/views
COPY --from=builder /go/src/github.com/lescactus/http-gallery-beego/static /app/static

EXPOSE 8080

RUN mkdir uploads

CMD ["./main"]
