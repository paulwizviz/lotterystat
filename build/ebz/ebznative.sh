#!/bin/sh

env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/package/macOS/${APP_NAME} ./cmd/ebz
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/package/linux/${APP_NAME} ./cmd/ebz
env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./build/package/windows/${APP_NAME}.exe ./cmd/ebz