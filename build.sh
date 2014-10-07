#!/usr/bin/env bashsh-0

CMD docker run --rm -v "$(pwd)":/helloworld -w /helloworld -e GOOS=windows -e GOARCH=amd64 golang:1.3-cross go build -v helloworld.go
