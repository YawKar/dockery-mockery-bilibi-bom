FROM golang:1.24.1-alpine3.21 AS base

COPY main.go main.go

EXPOSE 5050

RUN go build -o app main.go

ENTRYPOINT [ "./app" ]
