#!/bin/bash

if [ "$(basename $(realpath .))" != "lotterystat" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

export GO_VER=1.22.6-alpine3.19
export APP_NAME=ebz

export EBZ_BUILD_IMAGE=lotterystat/ebz-builder:current
export EBZ_BUILD_CONTAINER=ebz-container

COMMAND=$1
SUBCOMMAND=$2

# This is not use at the moment
function linuxBuild(){
    local cmd=$1
    case $cmd in
        "build")
            docker compose -f ./build/ebz/builder.yaml build
            docker compose -f ./build/ebz/builder.yaml up
            docker rm -f ${EBZ_BUILD_CONTAINER}
            docker rmi -f ${EBZ_IMAGE}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        "clean")
            docker rmi -f ${EBZ_IMAGE}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
    esac
}

function macBuild(){
    if [ ! -d ./package/macOS ]; then
        mkdir -p ./package/macOS
    fi
    env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o ./package/macOS/ebz ./cmd/ebz
}

function clean(){
    if [ -d ./package ]; then
        rm -rf ./package
    fi
    # linuxBuild clean
}

case $COMMAND in
    "build")
        macBuild
        # linuxBuild build
        ;;
    "clean")
        clean
        ;;
    *)
        echo "Usage: $0 [ build | clean ]"
        ;;
esac