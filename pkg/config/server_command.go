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

	config := pd.(*AppConfig)

	sub := tcl.InitInterp()
	sub.RegisterCommand("host", serverHostCommand, config)
	sub.RegisterCommand("port", serverPortCommand, config)
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
	config := pd.(*AppConfig)
	config.Host = argv[1]
	return "", nil
}

func serverPortCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "port", argv)
	}
	config := pd.(*AppConfig)
	config.Port, _ = strconv.Atoi(argv[1])
	return "", nil
}
