# Use a minimal base image.
FROM alpine:latest

# Install Hugo.
RUN ["apk", "add", "hugo"]

# Set the entrypoint to perform a Hugo build.
ENTRYPOINT ["hugo", \
    "--baseURL", "http://localhost:8080/", \
    "--source", "/site", \
    "--themesDir", "/themes", \
    "--theme", "hyde", \
    "--destination", "/public" \
]
