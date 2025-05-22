# syntax=docker/dockerfile:1.2
# Specifies the Dockerfile syntax version. This is good practice for newer features.

# --- Builder Stage ---
# Name this stage 'builder'
# This stage is responsible for compiling your Go application.
# Fault: Missing golang: prefix. Corrected.
FROM golang:1.24.3-alpine AS builder
# FROM 1.24.3-alpine AS builder # Original line with fault

# Install necessary build tools and packages for the Go compilation.
# apk update updates the package lists, apk add installs packages.
# --no-cache reduces image size by not storing the index locally after installation.
# build-base provides common compilation tools (GCC, make, etc.)
RUN apk update && apk add --no-cache build-base

# Set the working directory inside the builder image.
# All subsequent commands (COPY, RUN, etc.) will operate relative to this directory.
WORKDIR /build

# Copy the Go module definition files into the working directory.
# This allows Docker to cache the 'go mod download' step if these files don't change.
COPY go.mod go.sum ./

# Download Go module dependencies.
# This command fetches all required packages based on go.mod and go.sum.
RUN go mod download

# Build the Go application.
# -o /gaze-news: Specifies the output executable name and path (root directory of the build stage).
# ./cmd/http-server/main.go: Specifies the entry point for your Go application.
RUN go build -o /gaze-news ./cmd/http-server/main.go

# --- Runner Stage ---
# This stage creates a small, production-ready image containing only the compiled binary and its dependencies.
# It leverages Alpine Linux for its minimal size.
# Consider a specific Alpine version (e.g., alpine:3.20) for stability.
FROM alpine:latest
# alpine:latest can change, potentially leading to inconsistent builds over time.

# Set the working directory inside the final image.
WORKDIR /app

# Copy the compiled Go binary from the 'builder' stage into the 'runner' stage.
# --from=builder: Specifies to copy from the 'builder' stage.
# ./gaze-news: The path to the binary in the 'builder' stage.
# ./gaze-news: The destination path inside the current 'runner' stage's WORKDIR (/app).
# Fault: Missing leading slash for source path. Corrected.
COPY --from=builder /gaze-news ./gaze-news
# COPY --from=builder ./gaze-news ./gaze-news # Original line with fault

# Copy database migration files into the final image.
# This assumes your Go application directly uses these SQL files for migrations (e.g., via go:embed or direct path).
# Fault/Improvement: The destination path `/app/internal/adapters/db/sqlxdb/migrations/`
# implies your application will look for migrations at this specific path *inside the container*.
# If your app expects a top-level `migrations/` folder, this needs adjustment.
# Also, globbing `*.sql` is fine if you only have SQL files, but if you have other file types, refine it.
COPY ./internal/adapters/db/sqldb/migrations/*.sql /app/internal/adapters/db/sqlxdb/migrations/

# Expose the port on which your application listens.
# This merely documents the port; it doesn't publish it (that's done via `docker run -p` or `docker-compose.yml`).
EXPOSE 3030

# Define the default command to run when the container starts.
# Use the executable binary copied from the builder stage.
# The executable path is relative to the WORKDIR (`/app`).
CMD [ "./gaze-news" ]

# --- Explanatory Comments (moved from within the Dockerfile for clarity) ---
# Using Docker to create containers with the code binary, runtime system, and necessary tools.
# Docker containers are generally more efficient than traditional virtual machines because:
# 1. They share the host machine's operating system (OS) kernel, unlike VMs which run their own full OS copies.
# 2. Due to kernel sharing, their startup is quicker and they consume fewer resources (CPU, RAM).