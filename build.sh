#!/usr/bin/env bash
docker run -it --rm -v /home/rogeralsing/go:/app -e GOPATH="/app" -w "/app/src/github.com/rogeralsing/hellogo" golang sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hello'
docker build -t rogeralsing/hellogo .
docker run -p 8080:8080 rogeralsing/hellogo
