package config

import (
	"strconv"

	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func serverCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "server", argv)
	}

	sub := tcl.InitInterp()
	sub.RegisterCommand("host", serverHostCommand, pd.(*AppConfig))
	sub.RegisterCommand("port", serverPortCommand, pd.(*AppConfig))
	sub.RegisterCommand("env", envCommand, nil)
	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}
	return "", nil
}

func serverHostCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "host", argv)
	}
	pd.(*AppConfig).Host = argv[1]
	return "", nil
}

func serverPortCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "port", argv)
	}
	pd.(*AppConfig).Port, _ = strconv.Atoi(argv[1])
	return "", nil
}
