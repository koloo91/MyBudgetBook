package model

import (
	"time"
)

type ErrorVo struct {
	Message string `json:"message"`
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
		Id                   string
		Title                string
		Date                 time.Time
		Amount               float64
		CategoryId           string
		AccountId            string
		StandingOrderId      string
		StandingOrderPeriod  string
		StandingOrderLastDay time.Time
		Created              time.Time
		Updated              time.Time
	}

	BookingVo struct {
		Id                   string    `json:"id"`
		Title                string    `json:"title" binding:"required"`
		Date                 time.Time `json:"date" binding:"required"`
		Amount               float64   `json:"amount" binding:"required"`
		CategoryId           string    `json:"categoryId" binding:"required"`
		AccountId            string    `json:"accountId" binding:"required"`
		StandingOrderId      string    `json:"standingOrderId"`
		StandingOrderPeriod  string    `json:"standingOrderPeriod"`
		StandingOrderLastDay time.Time `json:"standingOrderLastDay"`
		Created              time.Time `json:"created"`
		Updated              time.Time `json:"updated"`
	}

	BookingsVo struct {
		Content []BookingVo `json:"content"`
	}
)

type (
	AccountBalance struct {
		AccountId string
		Name      string
		Balance   float64
	}

	AccountBalanceVo struct {
		AccountId string  `json:"accountId"`
		Name      string  `json:"name"`
		Balance   float64 `json:"balance"`
	}

	AccountBalancesVo struct {
		Content []AccountBalanceVo `json:"content"`
	}
)
