package models

type AdminModel struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Token    string `json:-`
}
