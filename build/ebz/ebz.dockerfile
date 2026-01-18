ARG GO_VER=1.22.6-alpine3.19

FROM golang:${GO_VER}

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go mod download


