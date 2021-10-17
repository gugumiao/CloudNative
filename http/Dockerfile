FROM golang:1.17-alpine AS builder
WORKDIR /http

COPY go.mod ./
COPY *.go ./
RUN go mod download \
    && go build -o httpserver .

FROM alpine:latest

COPY --from=builder /http/httpserver /

ENTRYPOINT ["/httpserver"]

