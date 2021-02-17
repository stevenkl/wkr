package models

type ExecutionParametersModel struct {
	Executable string `json:"executable"`
	Arguments []string `json:"arguments"`
}