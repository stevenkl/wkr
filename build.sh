#!/usr/bin/env sh

TARGET="wkr"

if [ $(uname -s) = "Windows_NT" ]; then
	TARGET="$TARGET.exe"
fi

SOURCES=main.go \
		config.go \
		config_commands.go \
		config_globals.go \
		config_validations.go

# # #
clear

go build -v -o $TARGET $SOURCES \
	&& ./wkr.exe
