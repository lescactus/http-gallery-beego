FROM library/golang:1.14.2-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags '-d -w -s' -o main


FROM alpine:3

WORKDIR /app

RUN chown -R 65534:65534 /app

COPY --from=builder --chown=65534:65534 /app/main /app
COPY --from=builder --chown=65534:65534 /app/views /app/views
COPY --from=builder --chown=65534:65534 /app/static /app/static

EXPOSE 8080

# nobody
USER 65534

CMD ["./main"]
