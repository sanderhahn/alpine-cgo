# Readme

Multistage dockerfile can be used to build minimal Alpine images with a Go binary.
The `golang-alpine-build` image has build tools and git so that we can build CGO binaries.
Added an SQLite query endpoint that requires CGO as an example (results in a 17.6MB docker image).

```bash
./build.sh

# run
docker run -p 8080:8080 -d --name hello -v `pwd`/data:/data hello
# open http://localhost:8080/

# shutdown
docker kill hello

# cleanup
docker system prune -f
```

## Development

```bash
go get github.com/codegangsta/gin

DATABASE=./data/sqlite.db gin run main.go
# open http://localhost:3000/
```
