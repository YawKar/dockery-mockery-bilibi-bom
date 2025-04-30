FROM golang:1.24.1-alpine3.21 AS base

COPY memory.go memory.go

RUN go build -gcflags '-N -l' -o memory memory.go

ENTRYPOINT [ "./memory" ]

CMD [ "100mb" ]
