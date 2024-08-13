#!/bin/sh

if [ "$(basename $(realpath .))" != "lotterystat" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

export PSQL_VER=16.2-alpine3.19
export PGADMIN_VER=8.9
export NETWORK=lotterystat_psql

COMMAND="$1"

case $COMMAND in
    "start")
        docker compose -f ./deployments/psql.yaml up
        ;;
    "stop")
        docker compose -f ./deployments/psql.yaml down
        ;;
    *)
        echo "Usage: $0 [start | stop]"
        ;;
esac