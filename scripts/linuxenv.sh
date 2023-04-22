#!/bin/bash

export OS_VER=ubuntu:22.10
export USER_NAME=ebzcli
export LINUXDT_IMAGE=paulwizviz/linuxdt:current

COMMAND=$1

function build(){
    docker-compose -f ./build/linuxenv/builder.yml build
}

function clean(){
    docker rmi -f ${LINUXDT_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function login(){
    docker run  -it --rm \
                --user ${USER_NAME}:${USER_NAME} \
                --workdir="/home/${USER_NAME}" \
                -v $PWD/build/ebenezer/package/linux/ebzcli:/usr/local/bin/ebzcli \
                ${LINUXDT_IMAGE} /bin/bash
}

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "login")
        login
        ;;
    *)
        echo "Usage: $0 [command]

command:
    build   linux desktop image
    clean   clear docker image caches
    login   to a linux environment"
        ;;
esac