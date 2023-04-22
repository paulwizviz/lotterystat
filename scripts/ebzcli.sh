#!/bin/bash

export GO_VER=golang:1.20.3
export APP_NAME=ebzcli

export EBZCLI_BUILD_IMAGE=paulwizviz/ebzcli:current
export EBZCLI_BUILD_CONTAINER=ebzcli_container

COMMAND=$1

function build(){
    docker-compose -f ./build/ebenezer/builder.yml build
    docker-compose -f ./build/ebenezer/builder.yml up
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