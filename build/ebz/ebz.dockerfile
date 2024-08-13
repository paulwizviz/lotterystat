ARG GO_VER

FROM golang:${GO_VER}

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

RUN apk add --no-cache libgit2 libgit2-dev git gcc-aarch64-none-elf g++ pkgconfig

