FROM golang:1.24.1-alpine3.21 AS base

COPY server.go server.go

RUN go build -o server server.go

EXPOSE 8080

ENTRYPOINT [ "./server" ]
