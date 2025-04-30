FROM golang:1.24.1-alpine3.21 AS base

COPY main.go main.go

RUN go build -o main main.go

ENTRYPOINT [ "./main" ]
