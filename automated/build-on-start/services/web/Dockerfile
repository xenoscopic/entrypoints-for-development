# Use a minimal Python image.
FROM python:alpine

# Install dependencies.
RUN ["pip3", "install", "flask"]

# Restart the server whenever the container starts.
ENTRYPOINT ["python", "/server/main.py"]
