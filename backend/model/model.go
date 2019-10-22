package model

import (
	"time"
)

type ErrorVo struct {
	Error string `json:"error"`
}

type (
	AccountVo struct {
		Id              string    `json:"id"`
		Name            string    `json:"name" binding:"required"`
		StartingBalance float64   `json:"startingBalance" binding:"required"`
		Created         time.Time `json:"created"`
		Updated         time.Time `json:"updated"`
	}

	Account struct {
		Id              string
		Name            string
		StartingBalance float64
		Created         time.Time
		Updated         time.Time
	}

	AccountsVo struct {
		Content []AccountVo `json:"content"`
	}
)

type (
	Category struct {
		Id       string
		ParentId string
		Name     string
		Created  time.Time
		Updated  time.Time
	}

	CategoryVo struct {
		Id       string    `json:"id"`
		ParentId string    `json:"parentId"`
		Name     string    `json:"name" binding:"required"`
		Created  time.Time `json:"created"`
		Updated  time.Time `json:"updated"`
	}

	CategoriesVo struct {
		Content []CategoryVo `json:"content"`
	}
)
