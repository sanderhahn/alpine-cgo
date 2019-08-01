# Readme

Multistage dockerfile can be used to build small Alpine images containing Go binaries.
The `golang-alpine-build` image has build tools and git so that we can build CGO projects.
Added a SQLite query as an example which results in a 17.6MB docker image.

## Docker

```bash
./build.sh

# run
docker run -p 8080:8080 -d --name alpine-cgo -v `pwd`/data:/data alpine-cgo
# open http://localhost:8080/

# shutdown
docker kill alpine-cgo

# cleanup
docker system prune -f
```

## Development

```bash
go get github.com/codegangsta/gin

DATABASE=./data/sqlite.db gin run main.go
# open http://localhost:3000/
```
