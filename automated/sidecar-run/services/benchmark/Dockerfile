# Use a minimal base image.
FROM alpine:latest

# Install dependencies.
RUN ["apk", "add", "apache2-utils"]

# Set the entrypoint to perform a benchmark of the web service.
ENTRYPOINT ["ab", "-n", "100", "-c", "10", "http://web/"]
