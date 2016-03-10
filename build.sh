#!/usr/bin/env bash
echo fixing godeps
${GOPATH}/bin/godep save
echo compiling Go via Docker
docker run -it --rm -v /home/rogeralsing/go:/app -e GOPATH="/app" -w "/app/src/github.com/rogeralsing/hellogo" golang sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hello'
echo Building Docker image
docker build -t rogeralsing/hellogo .
#docker run -p 8080:8080 rogeralsing/hellogo
docker push rogeralsing/hellogo

