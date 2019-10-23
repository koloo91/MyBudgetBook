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
		Id      string
		Name    string
		Created time.Time
		Updated time.Time
	}

	CategoryVo struct {
		Id      string    `json:"id"`
		Name    string    `json:"name" binding:"required"`
		Created time.Time `json:"created"`
		Updated time.Time `json:"updated"`
	}

	CategoriesVo struct {
		Content []CategoryVo `json:"content"`
	}
)

type (
	Booking struct {
		Id         string
		Title      string
		Comment    string
		Date       time.Time
		Amount     float64
		CategoryId string
		AccountId  string
		Created    time.Time
		Updated    time.Time
	}

	BookingVo struct {
		Id         string    `json:"id"`
		Title      string    `json:"title" binding:"required"`
		Comment    string    `json:"comment"`
		Date       time.Time `json:"date" binding:"required"`
		Amount     float64   `json:"amount" binding:"required"`
		CategoryId string    `json:"categoryId" binding:"required"`
		AccountId  string    `json:"accountId" binding:"required"`
		Created    time.Time `json:"created"`
		Updated    time.Time `json:"updated"`
	}

	BookingsVo struct {
		Content []BookingVo `json:"content"`
	}
)
