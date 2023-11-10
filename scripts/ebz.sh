#!/bin/bash

export GO_VER=golang:1.21-bullseye
export APP_NAME=ebz

export EBZ_BUILD_IMAGE=lotterystat/ebz:current
export EBZ_BUILD_CONTAINER=ebz_container

COMMAND=$1

function build(){
    # docker-compose -f ./build/ebenezer/builder.yml build
    # docker-compose -f ./build/ebenezer/builder.yml up
    go test -v ./...
    if [ "$(uname)" == "Darwin" ]; then
        echo "build for mac"
        env GOOS=darwin GOARCH=amd64 go build -o ./build/ebenezer/package/macOS/${APP_NAME} ./cmd/ebenezer/prod
    fi
    if [ "$(uname)" == "Linux" ]; then
        echo "build for linux"
        env GOOS=linux GOARCH=amd64 go build -o ./build/ebenezer/package/linux/${APP_NAME} ./cmd/ebenezer/prod
    fi
}

function clean(){
    docker rm -f ${EBZCLI_BUILD_CONTAINER}
    docker rmi -f ${EBZCLI_BUILD_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
    rm -rf ./build/ebenezer/package
}

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    *)
        echo "Usage: $0 [commands]

command:
    build  native cli app
    clean  remove containers, images and native packages"
        ;;
esac