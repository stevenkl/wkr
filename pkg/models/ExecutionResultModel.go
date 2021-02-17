package models

type ExecutionResultModel struct {
	StartedAt  string   `json:"started_at"`
	FinishedAt string   `json:"finished_at"`
	ExitCode   int      `json:"exit_code"`
	Log        []string `json:"log"`
}
