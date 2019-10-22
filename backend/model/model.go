package model

import (
	"time"
)

type AccountVo struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Account struct {
	Id      string
	Name    string
	Created time.Time
	Updated time.Time
}

type ErrorVo struct {
	Error string `json:"error"`
}

type AccountsVo struct {
	Content []AccountVo `json:"content"`
}
