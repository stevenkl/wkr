package config

import (
	"os"

	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)


func envCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "env", argv)
	}

	return os.Getenv(argv[1]), nil
}