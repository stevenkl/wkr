package main

import (
	"errors"
	
	"github.com/stevenkl/tcl.go/pkg/tcl"
)

// Config struct is the global Configuration element
type Config struct {
	Server  ServerConfig  `tcl:"ServerCommand"`
	Storage StorageConfig `tcl:"StorageCommand"`
	Users   UsersConfig   `tcl:"UsersCommand"`
	Jobs    JobsConfig    `tcl:"JobsCommand"`
}

type ServerConfig struct {
	Host string `tcl:"ServerHostCommand"`
	Port int    `tcl:"ServerPortCommand"`
}

type StorageConfig struct {
	Path string `tcl:"StoragePathCommand"`
}

type UserConfig struct {
	Name     string `tcl:"UserNameCommand"`
	Password string `tcl:"UserPasswordCommand"`
	Group    string `tcl:"UserGroupCommand"`
}

type UsersConfig struct {
	Users []*UserConfig
}

type JobConfig struct {
	Name string    `tcl:"JobNameCommand"`
	Workdir string `tcl:"JobWorkdirCommand"`
	Run string     `tcl:"JobRunCommand"`
}

type JobsConfig struct {
	Jobs []*JobConfig
}


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
	for _, user := range c.Users {
		if user.Name == name {
			return user, nil
		}
	}
	return nil, errors.New("not found")
}

func (c *UsersConfig) GetByGroup(group string) ([]*UserConfig, error) {
	users := make([]*UserConfig, 0)
	for _, user := range c.Users {
		if user.Group == group {
			users = append(users, user)
		}
	}
	return users, nil
}

func (j *JobConfig) Execute() error {
	// Trigger execution of job

	return nil
}

func (j *JobsConfig) GetByName(name string) (*JobConfig, error) {
	for _, job := range j.Jobs {
		if job.Name == name {
			return job, nil
		}
	}
	return nil, errors.New("not found")
}
