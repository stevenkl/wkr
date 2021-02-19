package main

import (
	"os"

	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func registerGlobalCommands(i *tcl.Interp) {
	i.RegisterCommand("env", envCommand, nil)
	i.RegisterCommand("if", cmds.If, nil)
}

func envCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "env", argv)
	}
	return os.Getenv(argv[1]), nil
}