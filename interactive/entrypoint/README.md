# Interactive `ENTRYPOINT`s

Interactive `ENTRYPOINT`s can be a really easy way to define development
workspaces or utility shells for performing specific tasks.


## Example Description

This example demonstrates creating an image with an interactive `ENTRYPOINT`
that has Python dependencies pre-installed.


## Usage

You can build the image using:

```
docker image build --tag analysis .
```

You can then create an interactive, ephemeral container with some analysis code
bind-mounted from the host side using:

```
docker run --interactive --tty --volume "${PWD}/analysis:/work/analysis" --rm analysis
```

This should drop you into an IPython shell. Try running some Python code:

```
from analysis.commits import load_commit_times
load_commit_times()
```

When you exit the shell, the container will be automatically removed. You can
remove the image with:

```
docker image rm analysis
```
