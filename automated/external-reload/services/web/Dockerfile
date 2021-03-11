# Use a minimal Python image.
FROM python:alpine

# Install dependencies.
RUN ["apk", "add", "nodejs", "npm"]
RUN ["npm", "install", "-g", "nodemon"]
RUN ["pip3", "install", "flask"]

# Use nodemon to watch code and automatically reload the server on changes.
ENTRYPOINT [ \
    "nodemon", "--watch", "/server", "/server/main.py" \
]
