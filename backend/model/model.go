package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
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

func (account *Account) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.New().String())
	scope.SetColumn("created", time.Now())
	scope.SetColumn("updated", time.Now())
	return nil
}

func (account *Account) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated", time.Now())
	return nil
}

type ErrorVo struct {
	Error string `json:"error"`
}
