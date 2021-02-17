package models

import "github.com/rs/xid"

type JobModel struct {
	ID            xid.ID                   `json:"id"`
	Title         string                   `json:"title"`
	CreatedAt     string                   `json:"created_at"`
	LastExecution string                   `json:"last_execution"`
	Executions    []string                 `json:"executions"`
	Parameters    ExecutionParametersModel `json:"params"`
}
