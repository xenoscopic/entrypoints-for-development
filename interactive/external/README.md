# Running Outside of Containers

When using Docker Desktop, it's possible to make network requests to the host
system using the special `host.docker.internal` DNS name. This can be useful in
cases where an application component isn't yet containerized, cases where you
need more direct access to a process (e.g. to use debugging tools outside of a
container), and cases where you just need to experiment with how a service is
configured.


## Example Description

This example demonstrates two services, one of which depends on the other. The
dependent service tries to contact the internal service on which it depends over
the `host.docker.internal` bridge.


## Usage

You'll need [Docker Desktop](https://www.docker.com/products/docker-desktop) to
run this demo (vanilla Docker doesn't support `host.docker.internal`,
[yet](https://github.com/docker/for-linux/issues/264)), as well as a
[Go](https://golang.org/) installation.

Open two terminal tabs and switch to this directory. In the first tab, run

```
docker-compose up --build
```

This will start a single web service (exposed on
[`http://localhost:8080`](http://localhost:8080)) that tries to contact another
backend service that it expects to be running on the host system. If you try to
access the web service before the backend service is running, it will return an
error message.

Now, in the second terminal tab, start the backend service on your local system
(i.e. not inside a container):

```
go run ./services/describer/describer.go
```

This will run the backend service on your local system on port 8081. This
backend service returns random descriptions that are used by the web frontend to
generate web pages. You can check that the backend service is running by
accessing it directly at [`http://localhost:8081`](http://localhost:8081).

Once it's running, go back to [`http://localhost:8080`](http://localhost:8080)
and you should see a message displayed successfully.

To shut down the project, simply `Ctrl+C` the `go` and `docker-compose` commands
and run

```
docker-compose down --rmi=local
```

back in the first terminal tab.
