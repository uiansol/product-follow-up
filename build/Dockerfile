# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app
ADD . /app/

RUN go build cmd/main.go

EXPOSE 8080

ENTRYPOINT ["./main"]