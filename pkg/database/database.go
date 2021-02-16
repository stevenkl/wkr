package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/stevenkl/wkr/pkg/config"
	"github.com/stevenkl/wkr/pkg/models"
)

var (
	ServerConfig config.AppConfig
	Instance     *Database
)

type Database struct {
	storage     string
	jobsPath    string
	resultsPath string
}

func Init(config *config.AppConfig) error {
	ServerConfig = *config
	Instance = new(Database)
	Instance.storage = ServerConfig.Data
	Instance.jobsPath = filepath.Join(Instance.storage, "jobs")
	Instance.resultsPath = filepath.Join(Instance.storage, "results")
	if err := makeFolders(); err != nil {
		return err
	}

	return nil
}

func (d *Database) Get(id string, job *models.JobModel) error {

	filePath := filepath.Join(Instance.jobsPath, id+".json")
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, job)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetAll() []models.JobModel {
	files, err := ioutil.ReadDir(d.jobsPath)
	if err != nil {
		return nil
	}
	jobs := make([]models.JobModel, 0)
	for _, f := range files {
		fmt.Println(f.Name())
		parts := strings.Split(f.Name(), ".")
		var job = models.JobModel{}
		err := d.Get(parts[0], &job)
		if err != nil {
			return nil
		}
		fmt.Println(job)
		jobs = append(jobs, job)
	}

	return nil
}

func makeFolders() error {
	err := os.MkdirAll(Instance.jobsPath, 0666)
	if err != nil {
		return err
	}
	err = os.MkdirAll(Instance.resultsPath, 0666)
	if err != nil {
		return err
	}
	return nil
}
