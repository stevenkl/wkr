package main

import (
	"errors"

	"github.com/rs/xid"
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Config struct is the global Configuration element
type Config struct {
	Server  ServerConfig  `tcl:"server,block"`
	Storage StorageConfig `tcl:"storage,block"`
	Users   UsersConfig   `tcl:"user,block"`
	Jobs    JobsConfig    `tcl:"job,block"`
}

type ServerConfig struct {
	Host string `tcl:"ServerHostCommand"`
	Port int    `tcl:"ServerPortCommand"`

	// Secret is generated at startup time
	Secret xid.ID
}

type StorageConfig struct {
	Path string `tcl:"StoragePathCommand"`
}

type UserConfig struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Group    string `json:"group"`
}


type UsersConfig []UserConfig

type JobConfig struct {
	ID      xid.ID `json:"id"`
	Name    string `json:"name"`
	Workdir string `josn:"workdir"`
	Run     string `json:"run"`
	LastRun int    `json:"last_run"`
}

type JobsConfig []JobConfig

// Parse parses the given string as a tcl script
func (c *Config) Parse(input string) error {
	i := tcl.InitInterp()

	// adding commands
	i.RegisterCommand("server", ServerCommand, nil)
	i.RegisterCommand("storage", StorageCommand, nil)
	i.RegisterCommand("user", UserCommand, nil)
	i.RegisterCommand("job", JobsCommand, nil)
	registerGlobalCommands(i)

	_, err := i.Eval(input)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Validate() error {

	var err error
	if err = validateServerConfig(c); err != nil {
		return err
	}
	if err = validateStorageConfig(c); err != nil {
		return err
	}
	if err = validateUsersConfig(c); err != nil {
		return err
	}
	if err = validateJobsConfig(c); err != nil {
		return err
	}
	return nil
}

// GetByName searches []*UsersConfig.users for a user with the given name
func (c *UsersConfig) GetByName(name string) (*UserConfig, error) {
	for _, user := range *c {
		if user.Name == name {
			return &user, nil
		}
	}
	return nil, errors.New("not found")
}

func (c *UsersConfig) GetByGroup(group string) ([]*UserConfig, error) {
	users := make([]*UserConfig, 0)
	for _, user := range *c {
		if user.Group == group {
			users = append(users, &user)
		}
	}
	return users, nil
}

func (j *JobConfig) Execute() error {
	// Trigger execution of job

	return nil
}

func (j *JobsConfig) GetByID(ID string) (*JobConfig, error) {
	for _, job := range *j {
		if job.ID.String() == ID {
			return &job, nil
		}
	}
	return nil, errors.New("not found")
}

func (j *JobsConfig) GetByName(name string) (*JobConfig, error) {
	for _, job := range *j {
		if job.Name == name {
			return &job, nil
		}
	}
	return nil, errors.New("not found")
}
