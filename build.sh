#!/usr/bin/env bash

GOOS=linux
GOARCH=amd64
# Binary
COMMANDS=(authorize cors login)
# Release
BUILD_DIR=./build

rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_DIR}

for CMDNAME in ${COMMANDS[*]}
do
    OUTPUT="${BUILD_DIR}/${CMDNAME}"
    # Build for linux
    echo "go build: ./cmd/${CMDNAME}.go"
    CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-w -s" -a -installsuffix cgo -o ${OUTPUT} ./cmd/${CMDNAME}.go
    # Write version
    echo "zip packing: ${CMDNAME}"
    (cd ${BUILD_DIR} && rm -f main && cp ${CMDNAME} main && zip aws-lambda-${CMDNAME}-${GOOS}-${GOARCH}.zip main)
done


