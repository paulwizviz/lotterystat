#!/bin/sh

#!/bin/sh

if [ "$(basename $(realpath .))" != "learn-sql" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

export PSQL_CLI_IMAGE=learn-sql/psqlcmd:current
export PSQL_CLI_CONTAINER=psqlcli_container
export NETWORK=learn-sql_psql

export PSQL_VER=16.2-alpine3.19
export PGADMIN_VER=8.5

export PSQL_VOL=psql_vol

COMMAND="$1"
SUBCOMMAND="$2"

function image(){
    local cmd="$1"
    case $cmd in
        "build")
            docker-compose -f ./build/psql/builder.yml build
            ;;
        "clean")
            docker rmi -f ${PSQL_CLI_IMAGE}
            docker volume rm ${PSQL_VOL}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        *)
            echo "image [ build | clean]"
            ;;
    esac
}

function client(){
    local cmd=$1
    case $cmd in
        "cli")
            docker run --network=${NETWORK} -v ${PWD}/db/psql/config/.pgpass:/root/.pgpass -v ${PWD}/db/psql/sql:/opt/sql/ -v ${PWD}/db/psql/scripts:/opt/scripts -w /opt -it --rm ${PSQL_CLI_IMAGE} /bin/bash
            ;;
        *)
            echo "client [ cli ]"
            ;;
    esac
}

function network(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployment/postgres/docker-compose.yml down
            ;;
        "start")
            docker-compose -f ./deployment/postgres/docker-compose.yml up
            ;;
        "stop")
            docker-compose -f ./deployment/postgres/docker-compose.yml down
            ;;
        *)
            echo "network [ clean | start | stop ]"
            ;;
    esac
}

case $COMMAND in
    "client")
        client $SUBCOMMAND
        ;;
    "clean")
        network stop
        network clean
        image clean
        ;;
    "image")
        image $SUBCOMMAND
        ;;
    "network")
        network $SUBCOMMAND
        ;;
    *)
        echo "$0 <commands>

commands:  
    clean    project setting
    image    build and clean
    network  clean, start and stop
"
        ;;
esac