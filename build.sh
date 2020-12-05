#!/bin/bash

function build_lib () {
    GOBIN=${PWD}/bin
    GOPATH=${GOPATH}:${PWD}
    go build github.com/Narasimha1997/httppool
}

build_lib