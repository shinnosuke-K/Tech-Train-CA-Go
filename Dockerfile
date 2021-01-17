FROM golang:1.15.6 AS builder
WORKDIR /go
COPY . /go
ENV GO111MODULE=on GOPATH=""
RUN CGO_ENABLED=0 GOOS=linux go build -o server server.go


FROM alpine:latest
RUN apk --no-cache add openssl
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && mkdir /root/rsa
WORKDIR /root/
COPY ./rsa ./rsa
COPY --from=builder /go/server .
