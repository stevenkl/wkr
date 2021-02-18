package config

import (
	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
	"github.com/stevenkl/wkr/pkg/models"
	"github.com/stevenkl/wkr/pkg/service"
)

func userCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "user", argv)
	}

	config := pd.(*AppConfig)
	user := models.UserModel{}

	sub := tcl.InitInterp()
	sub.RegisterCommand("name", userNameCommand, &user)
	sub.RegisterCommand("password", userPasswordCommand, &user)
	sub.RegisterCommand("group", userGroupCommand, &user)
	sub.RegisterCommand("env", envCommand, nil)

	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}
	config.Users = append(config.Users, user)
	return "", nil
}

func userNameCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "name", argv)
	}

	pd.(*models.UserModel).Name = argv[1]

	return "", nil
}

func userPasswordCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) < 2 || len(argv) > 3 {
		return "", cmds.ArityErr(i, "password", argv)
	}

	if argv[1] == "-raw" {
		// Hashing password
		hashed, err := service.HashPassword(argv[2])
		if err != nil {
			return "", err
		}
		pd.(*models.UserModel).PasswordHash = hashed

	} else {
		pd.(*models.UserModel).PasswordHash = argv[1]
	}

	return "", nil
}

func userGroupCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) >= 2 {
		pd.(*models.UserModel).Group = argv[1]
		return "", nil
	}
	pd.(*models.UserModel).Group = "guest"
	return "", nil
}