package config

import (
	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func adminCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "admin", argv)
	}

	sub := tcl.InitInterp()
	sub.RegisterCommand("user", adminUserCommand, pd.(*AppConfig))
	sub.RegisterCommand("password", adminPasswordCommand, pd.(*AppConfig))
	sub.RegisterCommand("env", envCommand, nil)

	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}
	return "", nil
}

func adminUserCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "user", argv)
	}

	pd.(*AppConfig).AdminUser = argv[1]

	return "", nil
}

func adminPasswordCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "password", argv)
	}

	pd.(*AppConfig).AdminPassword = argv[1]

	return "", nil
}