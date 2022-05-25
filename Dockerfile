# syntax=docker/dockerfile:1

# Setup container using alpine
FROM golang:1.16-alpine
WORKDIR /app

# Download go modules
COPY ./src/go.mod ./
COPY ./src/go.sum ./
RUN go mod download

# Copy source files and build
COPY ./src/. ./
RUN go build -o /go-image-service

# Execute server on 80
EXPOSE 80
ENV PORT=80
CMD ["/go-image-service"]