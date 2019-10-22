package model

import (
	"time"
)

type AccountVo struct {
	Id              string    `json:"id"`
	Name            string    `json:"name" binding:"required"`
	StartingBalance float64   `json:"startingBalance" binding:"required"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
}

type Account struct {
	Id              string
	Name            string
	StartingBalance float64
	Created         time.Time
	Updated         time.Time
}

type ErrorVo struct {
	Error string `json:"error"`
}

type AccountsVo struct {
	Content []AccountVo `json:"content"`
}
