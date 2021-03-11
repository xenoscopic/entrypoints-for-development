# Use a base image with Go support to perform our build.
FROM golang:alpine AS build

# Copy in source code and build the server.
COPY ["server.go", "template.html", "/"]
RUN ["go", "build", "-o", "/server", "/server.go"]

# Switch to a more minimal base image for deployment.
FROM alpine:latest

# Copy the server from the build stage.
COPY --from=build ["/server", "/usr/bin/web-server"]

# Set the entrypoint.
ENTRYPOINT ["/usr/bin/web-server"]
