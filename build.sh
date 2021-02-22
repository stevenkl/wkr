#!/usr/bin/env sh

OPTION=$1
TARGET="wkr"

if [ $(uname -s) = "Windows_NT" ]; then
	TARGET="$TARGET.exe"
fi

SOURCES=main.go \
		config.go \
		config_commands.go \
		config_globals.go \
		config_validations.go \
		web.go

# # #
clear

go build -v -o $TARGET $SOURCES


if [ "$OPTION" = "-run" ]; then
	"./$TARGET"
fi