FROM golang:1.24.1-alpine3.21 AS base

COPY cpu.go cpu.go

RUN go build -o cpu cpu.go

ENTRYPOINT [ "./cpu" ]

CMD [ "10" ]
