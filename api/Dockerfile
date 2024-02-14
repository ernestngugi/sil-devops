FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o sil-api cmd/sil-api/main.go

FROM alpine:3.18

RUN \
    apk update

COPY --from=builder /app/sil-api .

EXPOSE 3000

CMD ["./sil-api"]
