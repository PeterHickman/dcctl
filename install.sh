#!/bin/sh

BINARY='/usr/local/bin'

echo "Building dcctl"
go build dcctl.go

echo "Installing dcctl to $BINARY"
install -v dcctl $BINARY

echo "Removing the build"
rm dcctl
