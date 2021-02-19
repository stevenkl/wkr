package main

import (
	"strconv"
	
	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func ServerCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "server", argv)
	}

	sub := tcl.InitInterp()

	sub.RegisterCommand("host", ServerHostCommand, nil)
	sub.RegisterCommand("port", ServerPortCommand, nil)
	
	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}
	return "", nil
}

func ServerHostCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "host", argv)
	}
	config.Server.Host = argv[1]
	return "", nil
}

func ServerPortCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "port", argv)
	}
	config.Server.Port, _ = strconv.Atoi(argv[1])
	return "", nil
}


func StorageCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "storage", argv)
	}
	sub := tcl.InitInterp()
	sub.RegisterCommand("path", StoragePathCommand, nil)
	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}
	return "", nil
}

func StoragePathCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2  {
		return "", cmds.ArityErr(i, "path", argv)
	}
	config.Storage.Path = argv[1]
	return "", nil
}