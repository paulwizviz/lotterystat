#!/bin/bash

env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ./build/package/linux/${APP_NAME} ./cmd/ebenezer/prod
env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o ./build/package/macOS/${APP_NAME} ./cmd/ebenezer/prod
env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o ./build/package/windows/${APP_NAME}.exe ./cmd/ebenezer/prod