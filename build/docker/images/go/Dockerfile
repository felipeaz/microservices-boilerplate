FROM golang:alpine

# Add files to app dir
ADD . /app

# Set /app as the workdir
WORKDIR /app

# Install git & bash
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Update module dependencies
RUN go mod tidy