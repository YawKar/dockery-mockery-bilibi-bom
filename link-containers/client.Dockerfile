FROM golang:1.24.1-alpine3.21 AS base

COPY client.go client.go

RUN go build -o client client.go

ENTRYPOINT [ "./client" ]

CMD [ "server", "8080" ]
