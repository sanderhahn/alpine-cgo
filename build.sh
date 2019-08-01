#!/bin/bash
docker build -t golang-alpine-build golang-alpine-build
docker build -t alpine-cgo .
