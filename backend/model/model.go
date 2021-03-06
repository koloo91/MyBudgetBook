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
		IsMain          bool      `json:"isMain"`
		Created         time.Time `json:"created"`
		Updated         time.Time `json:"updated"`
	}

	Account struct {
		Id              string
		Name            string
		StartingBalance float64
		IsMain          bool
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

type (
	MonthStatistic struct {
		Expenses float64
		Incomes  float64
		Month    int
	}

	MonthStatisticVo struct {
		Expenses float64 `json:"expenses"`
		Incomes  float64 `json:"incomes"`
		Month    int     `json:"month"`
	}

	MonthStatisticsVo struct {
		Content []MonthStatisticVo `json:"content"`
	}
)

type (
	CategoryStatistic struct {
		Name string
		Sum  float64
	}

	CategoryStatisticVo struct {
		Name string  `json:"name"`
		Sum  float64 `json:"sum"`
	}

	CategoryStatisticsVo struct {
		Content []CategoryStatisticVo `json:"content"`
	}
)

type (
	InboxEntry struct {
		Id          string
		UserId      string
		BookingDate time.Time
		ValueDate   *time.Time
		IntendedUse string
		Amount      float64
		Created     time.Time
		Updated     time.Time
	}

	InboxEntryVo struct {
		Id          string     `json:"id"`
		BookingDate time.Time  `json:"bookingDate"`
		ValueDate   *time.Time `json:"valueDate"`
		IntendedUse string     `json:"intendedUse"`
		Amount      float64    `json:"amount"`
		Created     time.Time  `json:"created"`
		Updated     time.Time  `json:"updated"`
	}
)
