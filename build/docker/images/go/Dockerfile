FROM golang:alpine

WORKDIR /app

# Copy dir to workdir
COPY . /app

# Install git & bash
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Copy go mod and sum files
COPY go.mod go.sum /app/

# Update module dependencies
RUN go mod tidy