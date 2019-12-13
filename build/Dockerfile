FROM golang:1.13.4-alpine

RUN set -e \
    && apk add git \
    && mkdir -p /go/src/github.com/ieee0824 \
    && mkdir -p /tmp/build \
    && cd /go/src/github.com/ieee0824 \
    && git clone https://github.com/ieee0824/zip-line.git \
    && cd zip-line \
    && go build -o /bin/zipl cmd/zipl/main.go

WORKDIR /go/src/github.com/ieee0824/zip-line