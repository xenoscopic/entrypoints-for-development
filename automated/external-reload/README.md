# External Watch-Based Reload

If your framework or tooling doesn't provide automatic hot reloads (or its
functionality is flakey), then using an external tool to perform reloads on
filesystem changes can be a great alternative.


## Example Description

This example shows a Python Flask application that's reloaded on changes by an
external tool called [nodemon](https://github.com/remy/nodemon).


## Usage

To start the project, run:

```
docker-compose up --build
```

You can then access the site at
[`http://localhost:8080`](http://localhost:8080). If you modify the site source
code or content, then you'll see the `docker-compose up` output indicate that
the application has been rebuilt, and you can reload the application in your
browser.

When you're done, you can simply `Ctrl+C` the `docker-compose` command and then
tear down the project with:

```
docker-compose down --rmi=local
```
