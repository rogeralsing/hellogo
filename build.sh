#!/usr/bin/env bash
go build .
docker build -t rogeralsing/hellogo .
