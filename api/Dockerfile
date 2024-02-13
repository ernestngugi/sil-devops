FROM docker.io/golang:1.18-alpine as builder

WORKDIR /app

COPY . /app/

RUN \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o sil-api cmd/sil-api/main.go

FROM docker.io/alpine:3.14

RUN \
    apk update

COPY --from=builder /app/sil-api /usr/bin

EXPOSE 3000

CMD ["sil-api"]
