package models

type UserModel struct {
	Name         string `json:"name"`
	PasswordHash string `json:"password"`
	Group        string `json:"group"`
}
