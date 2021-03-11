# Automatic Hot Reload

Automatic hot reloads, if available with the tools or frameworks that you use,
provide an excellent mechanism for making `ENTRYPOINT`s both build and run code.


## Example Description

This example shows a Hugo site that's being automatically rebuilt and reloaded
on changes. Hugo actually pushes changes to your browser when it detects
updates, so you don't even need to refresh the site when you edit the source
code.


## Usage

To start the project, run:

```
docker-compose up --build
```

You can then access the site at
[`http://localhost:8080`](http://localhost:8080). If you modify the site source
code or content, then you'll see the `docker-compose up` output indicate that
the site has been rebuilt, and you'll see the site reloaded in your browser
automatically.

When you're done, you can simply `Ctrl+C` the `docker-compose` command and then
tear down the project with:

```
docker-compose down --rmi=local
```
