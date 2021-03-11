# `exec`ing into Containers

Containers aren't entirely closed off from the outside world. You can use the
`exec` commands of the Docker CLI and Docker Compose to spawn an interactive
shell inside of a container. This can be particularly useful for debugging, but
you can also use this technique to turn containers into reproducible development
environments.


## Example Description

This example contains a Go web server. Inside the image for the service,
debugging tools have also been installed, allowing you to attach to the
`ENTRYPOINT` of the container and debug it.

Because debugging requires special permissions, certain security features have
been disabled for the container. This is something you should only do in
development.


## Usage

Start the project with:

```
docker-compose up --build
```

You should be able to access the server on
[`http://localhost:8080`](http://localhost:8080).


Now try dropping into the container using:

```
docker-compose exec web /bin/sh
```

You'll be dropped into an interactive system shell inside the container. Start a
debugger and attach it to the `ENTRYPOINT` process. Then set a breakpoint and
let the process continue.

```
dlv attach 1
break handler /server.go:10
continue
```

Now try accessing the server again from your browser. If you hop back to the
terminal where you're running exec, you'll see that the break point has been
triggered. You'll have to type `continue` before the HTTP request will finish
being served.

This is just a basic example, but you can perform far more extensive debugging
operations with this same technique.

To shut down the project, simply `Ctrl+C` the `docker-compose up` command and
then run:

```
docker-compose down --rmi=local
```

This will automatically terminate your `exec` session.
