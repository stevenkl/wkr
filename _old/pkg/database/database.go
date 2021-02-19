package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/stevenkl/wkr/pkg/config"
	"github.com/stevenkl/wkr/pkg/models"
)

var (
	// ServerConfig is a link to package main.
	ServerConfig config.AppConfig
	// Instance of Database, accessible after calling database.Init()
	Instance *Database
)

// Database struct describes the (file-based) database
type Database struct {
	storage     string
	jobsPath    string
	resultsPath string
}

// Init database and creates an Instance for later access
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

// GetJob a single entry from the jobs database
func (d *Database) GetJob(id string, job *models.JobModel) error {

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

func (d *Database) SaveJob(job *models.JobModel) error {
	filePath := filepath.Join(d.jobsPath, job.ID.String()+".json")
	_, err1 := os.Stat(filePath)
	if err1 != nil {
		data, err2 := json.Marshal(job)
		if err2 != nil {
			return err2
		}
		err3 := ioutil.WriteFile(filePath, data, 0666)
		if err3 != nil {
			return err3
		}
		return nil
	}
	return errors.New("Job already exists")
}

func (d *Database) DeleteJob(job *models.JobModel) error {
	filePath := filepath.Join(d.jobsPath, job.ID.String()+".json")
	_, err1 := os.Stat(filePath)
	if err1 != nil {
		return err1
	}
	if err2 := os.Remove(filePath); err2 != nil {
		return err2
	}
	return nil

}

// GetAllJobs entries from the jobs database
func (d *Database) GetAllJobs() ([]*models.JobModel, error) {
	files, err := ioutil.ReadDir(d.jobsPath)
	if err != nil {
		return nil, err
	}
	jobs := make([]*models.JobModel, 0)
	for _, f := range files {
		parts := strings.Split(f.Name(), ".")
		var job = new(models.JobModel)
		err := d.GetJob(parts[0], job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (d *Database) GetExecutionResult(jobID string, runID int, res *models.ExecutionResultModel) error {
	filePath := filepath.Join(Instance.resultsPath, fmt.Sprintf("%s-%d", jobID, runID)+".json")
	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(content, res)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (d *Database) GetAllExecutionResults(jobID string) ([]*models.ExecutionResultModel, error) {
	job := new(models.JobModel)
	err := d.GetJob(jobID, job)
	if err != nil {
		return nil, err
	}
	executionResults := make([]*models.ExecutionResultModel, 0)
	for i := 1; i < job.Executions; i++ {
		res := new(models.ExecutionResultModel)
		err := d.GetExecutionResult(job.ID.String(), i, res)
		if err == nil {
			executionResults = append(executionResults, res)
		}
	}
	return executionResults, nil
}

// makeFolders for the database file structure
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
