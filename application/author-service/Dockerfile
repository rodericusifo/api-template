# Copyright 2025 Rodericus Ifo Krista
# SPDX-License-Identifier: MIT

###################################
# STEP 1: build executable binary #
###################################
FROM golang:alpine AS build-stage

# Set build env
ENV GOOS=linux
ENV CGO_ENABLED=0

# Install linux packages
RUN apk update
RUN apk add --no-cache git bash

# Setup working directory
WORKDIR /app

# Copy Go dependency files first
COPY go.mod go.sum ./

# Download and install all the dependencies
RUN go mod download
RUN go mod tidy
RUN go mod verify

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build Application
RUN go build -o /dist/server/main cmd/server/main.go

###############################
# STEP 2: build a small image #
###############################
FROM alpine:latest AS build-release-stage

# Set build arg
ARG MODE

# Validate build arg
RUN if [ -z "$MODE" ]; then echo "ERROR: MODE is required. Usage: docker build --build-arg MODE=dev ."; exit 1; fi

# Set build env
ENV MODE=${MODE}

# Setup working directory
WORKDIR /app

# Install for healthcheck in the release stage
RUN apk add --no-cache netcat-openbsd

# Copy the binary
COPY --from=build-stage /dist/server/main /dist/server/main

# Add a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Expose Port
EXPOSE 5050

# Use shell form for CMD
CMD ["sh", "-c", "/dist/server/main"]