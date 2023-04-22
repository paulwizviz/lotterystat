ARG GO_VER

FROM ${GO_VER}

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

