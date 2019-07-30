# Readme

Minimal Go server using multistage Docker to Alpine image (12Mb).

```bash
# build
docker build -t hello .

# run
docker run -p 8080:8080 -d --name hello hello
# open http://localhost:8080/world

# cleanup
docker system prune -f
```
