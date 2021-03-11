# Use a minimal base image.
FROM alpine:latest

# Install Hugo.
RUN ["apk", "add", "hugo"]

# Set Hugo as the container entrypoint and automatically rebuild/reload content
# whenever changes are detected.
ENTRYPOINT ["hugo", "server", \
    "--bind", "0.0.0.0", \
    "--source", "/site", \
    "--themesDir", "/themes", \
    "--theme", "hyde" \
]
