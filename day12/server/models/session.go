package models

type Session struct {
	Username string `json:"-"`
	Token    string `json:"-"`
}
