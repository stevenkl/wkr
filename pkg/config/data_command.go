package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/stevenkl/tcl.go/pkg/cmds"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

func dataCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "data", argv)
	}

	sub := tcl.InitInterp()
	sub.RegisterCommand("path", dataPathCommand, pd.(*AppConfig))
	sub.RegisterCommand("env", envCommand, nil)
	_, err := sub.Eval(argv[1])
	if err != nil {
		return "", err
	}
	return "", nil
}

func dataPathCommand(i *tcl.Interp, argv []string, pd interface{}) (string, error) {
	if len(argv) != 2 {
		return "", cmds.ArityErr(i, "data", argv)
	}

	dataPath := argv[1]
	if !filepath.IsAbs(dataPath) {
		pwd, _ := os.Getwd()
		dataPath, _ = filepath.Abs(filepath.Join(pwd, dataPath))
	}

	stat, err := os.Stat(dataPath)
	if err != nil {
		return "", err
	}
	if !stat.IsDir() {
		return "", errors.New("Data path is not a directory")
	}

	pd.(*AppConfig).Data = dataPath
	return "", nil
}
