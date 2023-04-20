# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /go/src/form3-http-go

COPY go.mod go.sum ./
RUN go mod download

COPY . /go/src/form3-http-go

ENTRYPOINT ["./entrypoint.sh"]