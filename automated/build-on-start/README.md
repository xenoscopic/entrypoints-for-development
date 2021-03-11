# Re-building/Re-Running on Container (Re)Start

Re-building (or reloading) and then re-running code in sync with your container
starting (or restarting) can be an easy way to manually control when code is
reloaded/re-built. This works well with interpreted languages (like Python or
Ruby), and also with tools that provide a way to build and run code in one step
(like Go's `go run` or a Gulp-based build script). But even if your toolchain
doesn't provide this functionality, you can emulate it with a script that
performs a combined build + run operation.


## Example Description

This example contains a simple Flask application that's reloaded whenever a
container starts or restarts.


## Usage


To start the project, run:

```
docker-compose up
```

You can then access the server at
[`http://localhost:8080`](http://localhost:8080).

If you update the code, you can then reload the server by `Ctrl-C`ing the
`docker-compose up` command and then just run:

```
docker-compose up
```

In the normal case, where you have multiple services, you could just run:

```
docker-compose restart web
```

In this example, with only one service, doing this will cause the initial `up`
command to terminate when the `web` service is recreated because there are no
other services running. In a project with multiple services, the original `up`
command would continue to function and would attach to the output of the
recreated container. You can test this here by adding an additional service
(e.g. a Redis service) to the project. In any case, if you're only dealing with
one service, then you're better off simply using `Ctrl-C` and then restarting
the `docker-compose up --build` command anyway since it's less work.

To shut down the example, simply `Ctrl-C` the `docker-compose up` instance and
then tear down the project with:

```
docker-compose down --rmi=local
```
