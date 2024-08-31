FROM golang:1.22.4-alpine AS builder

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy
