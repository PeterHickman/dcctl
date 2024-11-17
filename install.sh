#!/bin/sh

BINARY='/usr/local/bin'
APP=dcctl

echo "Building $APP"
go build -ldflags="-s -w" $APP.go

echo "Installing $APP to $BINARY"
install -v $APP $BINARY

echo "Removing the build"
rm $APP
