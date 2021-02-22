package main

import (
	"strconv"

	"github.com/rs/xid"
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
	registerGlobalCommands(sub)
	
	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}
	config.Server.Secret = xid.New()
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
	registerGlobalCommands(sub)
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


func UserCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "user", argv)
	}

	user := new(UserConfig)
	
	sub := tcl.InitInterp()
	sub.RegisterCommand("name", UserNameCommand, user)
	sub.RegisterCommand("password", UserPasswordCommand, user)
	sub.RegisterCommand("group", UserGroupCommand, user)
	registerGlobalCommands(sub)

	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}

	config.Users = append(config.Users, *user)
	return "", nil
}

func UserNameCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "name", argv)
	}
	user := pd.(*UserConfig)
	user.Name = argv[1]
	return "", nil
}

func UserPasswordCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) < 2 || len(argv) > 3 {
		return "", cmds.ArityErr(i, "password", argv)
	}
	user := pd.(*UserConfig)

	if argv[1] == "-raw" {
		var err error
		user.Password, err = generatePasswordHash(argv[2])
		if err != nil {
			return "", err
		}
		return "", nil
	}
	user.Password = argv[1]
	return "", nil
}

func UserGroupCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "group", argv)
	}
	user := pd.(*UserConfig)
	if argv[1] == "" {
		user.Group = "guest"
		return "", nil
	}
	user.Group = argv[1]
	return "", nil
}


func JobsCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "job", argv)
	}

	job := new(JobConfig)
	sub := tcl.InitInterp()
	sub.RegisterCommand("id", JobIDCommand, job)
	sub.RegisterCommand("name", JobNameCommand, job)
	sub.RegisterCommand("workdir", JobWorkdirCommand, job)
	sub.RegisterCommand("run", JobRunCommand, job)
	registerGlobalCommands(sub)

	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}

	config.Jobs = append(config.Jobs, *job)
	return "", nil
}

func JobIDCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "id", argv)
	}
	job := pd.(*JobConfig)
	guid, err := xid.FromString(argv[1])
	if err != nil {
		job.ID = xid.New()
	}
	job.ID = guid
	return "", nil
}

func JobNameCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "name", argv)
	}
	job := pd.(*JobConfig)
	job.Name = argv[1]
	return "", nil
}

func JobWorkdirCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "workdir", argv)
	}
	job := pd.(*JobConfig)
	job.Workdir = argv[1]
	return "", nil
}

func JobRunCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "run", argv)
	}
	job := pd.(*JobConfig)
	job.Run = argv[1]
	return "", nil
}
