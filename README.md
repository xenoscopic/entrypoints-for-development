# Container `ENTRYPOINT` Patterns for Development

This project contains examples of container `ENTRYPOINT` patterns for
development. It's certainly not an exhaustive list, so feel free to fork and
hack together new ideas. To run these examples, you'll need Docker Compose and
some sort of Docker engine. You can get both easily by installing
[Docker Desktop](https://www.docker.com/products/docker-desktop) for macOS or
Windows.

This project grew out of a
[talk](https://havoc.io/talks/entrypoints-for-development) given at the Docker
Community All Hands meeting on 11 March 2021.


## Cloning

This repository uses Git submodules to facilitate demos using external projects.
To ensure you get a copy of these submodules when cloning, use:

```
git clone --recurse-submodules https://github.com/havoc-io/entrypoints-for-development.git
```

If you've already clone the repository, no worries, just run:

```
git submodule update --init --recursive
```


## Automated Development Patterns with Containers

The `automated` directory contains examples of recompiling and rebuilding
applications and their assets inside containers using a variety of techniques:

- Automatic hot reload ([automated/hot-reload](automated/hot-reload))
- External watch-based reload ([automated/external-reload](automated/external-reload))
- Re-building/re-running on container (re)start ([automated/build-on-start](automated/build-on-start))
- Manual sidecar runs ([automated/sidecar-run](automated/sidecar-run))
- Full image rebuilds ([automated/image-rebuild](automated/image-rebuild))


## Interactive Development Patterns with Containers

The `interactive` directory contains examples of working interactively in
containers:

- `exec`ing into containers ([interactive/exec](interactive/exec))
- Interactive `ENTRYPOINT`s ([interactive/entrypoint](interactive/entrypoint))
- Running outside of containers ([interactive/external](interactive/external))
