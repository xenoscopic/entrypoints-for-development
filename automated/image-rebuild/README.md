# Full Image Rebuilds

Full image rebuilds are the slowest (and heaviest) way to rebuild services in a
Compose project, but they are perhaps the most flexible and reproducible. They
are most useful for code that needs to be recompiled, especially if that code
needs to be compiled as part of the first stage of a multi-stage build.


## Example Description

This example shows a case where a full image rebuild might be necessary: a Go
server that's compiled as part of a multi-stage build. This particular server
also embeds static resources (specifically a page template) that are only
available during the build stage.


## Usage

To use this example, simply run

```
docker-compose up --build
```

You can then access the server at
[`http://localhost:8080`](http://localhost:8080).

When you want to rebuild this service, you can `Ctrl+C` the `docker-compose`
command and then just re-run it to rebuild the service.

If you have multiple services and only want to rebuild one (without restarting
the whole project), you can switch to a separate terminal tab and do:

```
docker-compose build web
docker-compose up --detach web
```

Note that in this example, with only one service, doing this will cause the
initial `up` command to terminate when the `web` service is recreated because
there are no other services running. In a project with multiple services, the
original `up` command would continue to function and would attach to the output
of the recreated container. You can test this here by adding an additional
service (e.g. a Redis service) to the project. In any case, if you're only
dealing with one service, then you're better off simply using `Ctrl-C` and then
restarting the `docker-compose up --build` command anyway since it's less work.

To shut down the example, simply run:

```
docker-compose down --rmi=local
```
