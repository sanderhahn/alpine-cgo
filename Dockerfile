FROM golang-alpine-build AS builder
WORKDIR /go/src/alpine-cgo
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine
ENV PORT 8080
ENV DATABASE /data/sqlite.db
ENV PATH $PATH:/go/bin/
EXPOSE $PORT
COPY --from=builder /go/bin/alpine-cgo /go/bin/alpine-cgo
ENTRYPOINT ["alpine-cgo"]
