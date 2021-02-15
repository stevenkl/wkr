package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stevenkl/tcl.go/pkg/tcl"
)

type AppConfig struct {
	Host          string
	Port          int
	Data          string
	AdminUser     string
	AdminPassword string
}

func New() (AppConfig, error) {
	appcfg := AppConfig{}

	pwd, _ := os.Getwd()
	cfgFile := filepath.Join(pwd, "wkr.config")
	cfgFileContent, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return appcfg, err
	}

	i := tcl.InitInterp()
	i.RegisterCommand("server", serverCommand, &appcfg)
	i.RegisterCommand("data", dataCommand, &appcfg)
	i.RegisterCommand("admin", adminCommand, &appcfg)

	_, errr := i.Eval(string(cfgFileContent))
	if errr != nil {
		return appcfg, errr
	}

	return appcfg, nil
}
