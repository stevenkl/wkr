#!/usr/bin/env sh

TARGETDIR=./build

TARGETNAME=wkr
if [ $(uname) = "Windows_NT" ]; then
	TARGETNAME="${TARGETNAME}.exe"
fi

function go_build() {
	go build -o "$TARGETDIR/$TARGETNAME" "./cmd/wkr"
}


echo "Building $TARGETDIR/$TARGETNAME"
go_build
