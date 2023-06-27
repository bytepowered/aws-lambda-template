#!/usr/bin/env bash

GOOS=linux
GOARCH=amd64
# Binary
COMMANDS=(login)
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
    (cd ${BUILD_DIR} && rm -f main && cp ${CMDNAME} main && zip -9 aws-lambda-${CMDNAME}-${GOOS}-${GOARCH}.zip main)
done

ENVFILES=(env)
for ENVF in ${ENVFILES[*]}
do
    echo "zip env file: aws-${ENVF}-layer.zip"
    (cd conf.d && zip aws-${ENVF}-layer.zip ".${ENVF}")
done

