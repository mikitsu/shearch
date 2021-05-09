FROM golang:latest as builder

WORKDIR /tmp/build
COPY main.go go.mod .
COPY lib lib
RUN CGO_ENABLED=0 go build .

FROM alpine:latest

COPY --from=builder /tmp/build/shearch /bin
