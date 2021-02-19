package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stevenkl/tcl.go/pkg/tcl"
	"github.com/stevenkl/wkr/pkg/models"
)

// AppConfig represents the global configuration for Wkr
type AppConfig struct {
	Host  string
	Port  int
	Data  string
	Users []models.UserModel
}

// New returns a newly AppConfig
func New() AppConfig {
	appcfg := new(AppConfig)

	return *appcfg
}

// Parse parses the given string as filepath for configuration
func (cfg *AppConfig) Parse(filePath string) error {

	pwd, _ := os.Getwd()
	cfgFile := filepath.Join(pwd, filePath)
	cfgFileContent, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return err
	}

	i := tcl.InitInterp()
	i.RegisterCommand("server", serverCommand, cfg)
	i.RegisterCommand("storage", storageCommand, cfg)
	i.RegisterCommand("user", userCommand, cfg)

	_, errr := i.Eval(string(cfgFileContent))
	if errr != nil {
		return errr
	}

	return nil
}
