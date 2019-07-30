FROM golang:latest AS builder
WORKDIR /go/src/hello
COPY . .
# ENV CGO_ENABLED 0
# ENV GOOS linux
# ENV GOARCH amd64
RUN go get -d -v ./...
# RUN go install -v ./...
RUN go build -o /go/bin/hello
RUN ldd /go/bin/hello

FROM alpine
ENV PORT 8080
ENV PATH $PATH:/go/bin/
ENV LD_LIBRARY_PATH /lib
EXPOSE $PORT
COPY --from=builder /go/bin/hello /go/bin/hello
# alpine has muslc and go is compiled with glibc, but they are compatible
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENTRYPOINT ["hello"]
