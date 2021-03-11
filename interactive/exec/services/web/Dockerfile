# Use a minimal base image with Go support.
FROM golang:alpine AS build

# Install dependencies.
RUN ["go", "install", "github.com/go-delve/delve/cmd/dlv@latest"]

# Copy in source code and build the server.
COPY ["server.go", "/server.go"]
RUN ["go", "build", "-o", "/server", "/server.go"]

# Set the entrypoint.
ENTRYPOINT ["/server"]
