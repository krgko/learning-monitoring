# Server

Using golang with echo framework to build API server

## Multistage build

```bash
## Builder Image ##
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

## Outcome Image ##
FROM alpine:latest

COPY --from=builder /build/server .

CMD ["./server"]
```

Then, what is `CGO_ENABLED` and `GOOS` while building the source code

- `CGO_ENABLED` is for building static binary and set `0` is for breaking linkage between C System Library for our standalone binary file
- `GOOS` is for setting OS as linux or mac or windows

## Reference Document

- [echo-customization](https://echo.labstack.com/guide/customization)
