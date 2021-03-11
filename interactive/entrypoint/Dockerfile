# Start with a base image that's suitable for doing some analysis work.
FROM amancevice/pandas:slim

# Install additional dependencies. You'd normally use a requirements.txt file to
# specify these, but for demonstration purposes this is fine.
RUN ["pip3", "install", "ipython", "requests"]

# Create a working directory where we can run analyses. We'll mount our analysis
# code into this using a bind mount.
RUN ["mkdir", "/work"]
WORKDIR "/work"

# Use an IPython shell as the entrypoint.
ENTRYPOINT ["ipython3"]
