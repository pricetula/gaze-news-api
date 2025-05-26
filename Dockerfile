# syntax=docker/dockerfile:1.2

# --- Builder Stage ---
FROM golang:1.24.3-alpine AS builder

# create a working directory for the Go application
# This is where the Go source code will be copied and built.
WORKDIR /app

# copy the Go module files to the working directory
# This allows Docker to cache the module download step, speeding up builds.
COPY go.mod go.sum ./

# Download the Go module dependencies.
RUN go mod download

# Copy the entire application source code into the working directory.
# This includes all Go files, static assets, and other necessary files.
# It's important to do this after downloading the modules to leverage Docker's caching.
COPY . .

# Build the Go application:
# CGO_ENABLED=0: Critical for creating a truly static binary, allowing tiny base images.
# GOOS=linux: Ensures the binary is compiled for Linux.
# -o main: Output the executable as 'main' in the current WORKDIR (/app).
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/http-server/main.go

# --- Runner Stage ---
# For a Go app with CGO_ENABLED=0, distroless/static-debian11 is a great choice.
# It includes CA certificates for HTTPS calls but no shell or package manager.
FROM gcr.io/distroless/static-debian11

# If you need basic OS utilities or a shell for debugging, alpine:3.20 is a good alternative.
# FROM alpine:3.20

WORKDIR /app

# Copy the compiled Go binary from the 'builder' stage into the 'runner' stage.
# It's now located at /app/main in the builder.
COPY --from=builder /app/main ./main

# Copy database migration files.
# IMPORTANT: Verify 'sqldb' vs 'sqlxdb' path.
# Also, ensure your Go app expects to find migrations at this exact path '/app/internal/adapters/db/sqlxdb/migrations/'
# If your app uses go:embed, you don't need this COPY.
COPY ./migrations/*.sql /app/migrations/

# Expose the port your application listens on.
EXPOSE 3030

# Define the default command to run when the container starts.
# The executable path is relative to the WORKDIR (`/app`).
CMD [ "./main" ]
