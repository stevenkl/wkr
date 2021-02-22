package main

import (
	"errors"
	"os"
)

func validateServerConfig(c *Config) error {
	if c.Server.Host == "" {
		return errors.New("server.host field must not an empty string")
	}
	if c.Server.Port == 0 {
		return errors.New("server.port field must greater than 0")
	}
	return nil
}

func validateStorageConfig(c *Config) error {
	_, err := os.Stat(c.Storage.Path)
	if err != nil {
		return err
	}
	return nil
}

func validateUsersConfig(c *Config) error {
	if len(c.Users) <= 0 {
		return errors.New("it must at least one user be defined")
	}
	return nil
}

func validateJobsConfig(c *Config) error {
	if len(c.Jobs) <= 0 {
		return errors.New("it must at least one job be defined")
	}
	return nil
}
