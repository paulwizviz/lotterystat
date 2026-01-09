ARG GO_VER=1.25

FROM golang:${GO_VER}

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go mod download


