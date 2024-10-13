FROM golang:1.23.1
# FROM golang:1.23.1-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY main.go go.* /app/