FROM golang:1.22 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags '-d -w -s' -o main

FROM scratch

WORKDIR /app

COPY --from=builder /app/main /
COPY --from=builder --chown=65534:65534 /app/main /app
COPY --chown=65534:65534 ./views /app/views
COPY --chown=65534:65534 ./static /app/static
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["/app/main"]