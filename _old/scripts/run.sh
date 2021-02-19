#! /usr/bin/env sh

TARGETDIR=./build

TARGETNAME=wkr
if [ $(uname) = "Windows_NT" ]; then
	TARGETNAME="${TARGETNAME}.exe"
fi

exec "$TARGETDIR/$TARGETNAME"