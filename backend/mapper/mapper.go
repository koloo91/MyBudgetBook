package mapper

import (
	"github.com/google/uuid"
	"github.com/koloo91/model"
	"time"
)

func AccountVoToEntity(vo model.AccountVo) model.Account {
	return model.Account{
		Id:              uuid.New().String(),
		Name:            vo.Name,
		StartingBalance: vo.StartingBalance,
		IsMain:          vo.IsMain,
		Created:         time.Now(),
		Updated:         time.Now(),
	}
}

func AccountEntityToVo(entity model.Account) model.AccountVo {
	return model.AccountVo{
		Id:              entity.Id,
		Name:            entity.Name,
		StartingBalance: entity.StartingBalance,
		IsMain:          entity.IsMain,
		Created:         entity.Created,
		Updated:         entity.Updated,
	}
}

func AccountEntitiesToVos(entities []model.Account) []model.AccountVo {
	vos := make([]model.AccountVo, 0, len(entities))
	for _, entity := range entities {
		vos = append(vos, AccountEntityToVo(entity))
	}
	return vos
}

func CategoryVoToEntity(vo model.CategoryVo) model.Category {
	return model.Category{
		Id:      uuid.New().String(),
		Name:    vo.Name,
		Created: time.Now(),
		Updated: time.Now(),
	}
}

func CategoryEntityToVo(entity model.Category) model.CategoryVo {
	return model.CategoryVo{
		Id:      entity.Id,
		Name:    entity.Name,
		Created: entity.Created,
		Updated: entity.Updated,
	}
}

func CategoryEntitiesToVos(entities []model.Category) []model.CategoryVo {
	vos := make([]model.CategoryVo, 0, len(entities))
	for _, entity := range entities {
		vos = append(vos, CategoryEntityToVo(entity))
	}
	return vos
}

func BookingVoToEntity(vo model.BookingVo) model.Booking {
	return model.Booking{
		Id:                   uuid.New().String(),
		Title:                vo.Title,
		Date:                 vo.Date,
		Amount:               vo.Amount,
		CategoryId:           vo.CategoryId,
		AccountId:            vo.AccountId,
		StandingOrderId:      vo.StandingOrderId,
		StandingOrderPeriod:  vo.StandingOrderPeriod,
		StandingOrderLastDay: vo.StandingOrderLastDay,
		Created:              time.Now(),
		Updated:              time.Now(),
	}
}

func BookingEntityToVo(entity model.Booking) model.BookingVo {
	return model.BookingVo{
		Id:                   entity.Id,
		Title:                entity.Title,
		Date:                 entity.Date,
		Amount:               entity.Amount,
		CategoryId:           entity.CategoryId,
		AccountId:            entity.AccountId,
		StandingOrderId:      entity.StandingOrderId,
		StandingOrderPeriod:  entity.StandingOrderPeriod,
		StandingOrderLastDay: entity.StandingOrderLastDay,
		Created:              entity.Created,
		Updated:              entity.Updated,
	}
}

func BookingEntitiesToVos(entities []model.Booking) []model.BookingVo {
	vos := make([]model.BookingVo, 0, len(entities))
	for _, entity := range entities {
		vos = append(vos, BookingEntityToVo(entity))
	}
	return vos
}

func AccountBalanceEntityToVo(entity model.AccountBalance) model.AccountBalanceVo {
	return model.AccountBalanceVo{
		AccountId: entity.AccountId,
		Name:      entity.Name,
		Balance:   entity.Balance,
	}
}

func AccountBalanceEntitiesToVos(entities []model.AccountBalance) []model.AccountBalanceVo {
	vos := make([]model.AccountBalanceVo, 0, len(entities))
	for _, entity := range entities {
		vos = append(vos, AccountBalanceEntityToVo(entity))
	}
	return vos
}

func MonthStatisticEntityToVo(entity model.MonthStatistic) model.MonthStatisticVo {
	return model.MonthStatisticVo{
		Expenses: entity.Expenses,
		Incomes:  entity.Incomes,
		Month:    entity.Month,
	}
}

func MonthStatisticEntitiesToVos(entities []model.MonthStatistic) []model.MonthStatisticVo {
	vos := make([]model.MonthStatisticVo, 0, len(entities))
	for _, entity := range entities {
		vos = append(vos, MonthStatisticEntityToVo(entity))
	}
	return vos
}

func CategoryStatisticEntityToVo(entity model.CategoryStatistic) model.CategoryStatisticVo {
	return model.CategoryStatisticVo{
		Name: entity.Name,
		Sum:  entity.Sum,
	}
}

func CategoryStatisticEntitiesToVos(entities []model.CategoryStatistic) []model.CategoryStatisticVo {
	vos := make([]model.CategoryStatisticVo, 0, len(entities))
	for _, entity := range entities {
		vos = append(vos, CategoryStatisticEntityToVo(entity))
	}
	return vos
}
