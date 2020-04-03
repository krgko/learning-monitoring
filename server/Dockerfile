FROM golang:1.13.7-alpine3.11 AS builder
RUN mkdir -p /build/src

RUN apk update \
    && apk add --update git \
    && apk add --update openssh

ENV GOPATH=/build
ENV GOBIN=/usr/local/go/bin
ENV GO111MODULE=on

WORKDIR $GOPATH/src

ADD . .
RUN go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -a -o server . \
    && cp server $GOPATH/.

FROM alpine:latest

COPY --from=builder /build/server .

CMD ["./server"]
