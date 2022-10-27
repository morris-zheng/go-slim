FROM golang:1.19 as base

LABEL stage=base

ENV GOPROXY="https://goproxy.cn"

WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

FROM alpine:3.12

WORKDIR /app

COPY --from=base /app/main /app/main

ENTRYPOINT ["/app/main"]
EXPOSE 3000
