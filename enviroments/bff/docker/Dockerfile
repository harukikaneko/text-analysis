FROM golang:1.19.0-alpine3.16 AS builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ .
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN go test ./...

RUN go build -o target/bff

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/target/bff .

CMD ["./bff"]