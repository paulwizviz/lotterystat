#!/bin/bash

env GOOS=linux GOARCH=amd64 go build -o ./build/package/linux/${APP_NAME} ./cmd/ebenezer/ebzcli
env GOOS=darwin GOARCH=amd64 go build -o ./build/package/macOS/${APP_NAME} ./cmd/ebenezer/ebzcli
env GOOS=windows GOARCH=amd64 go build -o ./build/package/windows/${APP_NAME}.exe ./cmd/ebenezer/ebzcli