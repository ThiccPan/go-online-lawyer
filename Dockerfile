#syntax=docker/dockerfile:1

# stage 1
FROM golang:1.20 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build app/main.go
EXPOSE 8080
CMD ["/app/main"]