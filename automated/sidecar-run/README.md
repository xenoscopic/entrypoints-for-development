# Manual Sidecar Runs

Sidecar services that aren't part of your normal project, but which you run on
an *ad hoc* basis via `docker-compose run`, can be a great way to automate tasks
or replace utility scripts.


## Example Description

This example shows a Hugo site that's built by a sidecar container and served by
a separate Go server in another container. The site is built and shared in a
Docker volume that's connected to both services.


## Usage

Before starting the project, perform an initial build using the sidecar service
using:

```
docker-compose run --rm hugo
```

You can then start the primary server with:

```
docker-compose up --build
```

You can then access the site at
[`http://localhost:8080`](http://localhost:8080).

If you make changes to the site code, you can then rebuild the site by running
the sidecar service again:

```
docker-compose run --rm hugo
```

and then refresh the site in your browser.

When you're done, you can simply `Ctrl+C` the `docker-compose up` command and
then tear down the project with:

```
docker-compose down --rmi=local --volumes
```
