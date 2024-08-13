#!/bin/sh

env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -tags static,system_libgit2 -o ./build/package/linux/${APP_NAME} ./cmd/ebz
# env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -tags static,system_libgit2 -o ./build/package/windows/${APP_NAME}.exe ./cmd/ebz