# Flask imports
from flask import Flask

# Create an application.
app = Flask("demo")

# Define a simple handler.
@app.route('/')
def index():
    return "Hello there!"

# Run the application.
app.run("0.0.0.0")
