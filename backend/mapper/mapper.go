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
		Created:         time.Now(),
		Updated:         time.Now(),
	}
}

func AccountEntityToVo(entity model.Account) model.AccountVo {
	return model.AccountVo{
		Id:              entity.Id,
		Name:            entity.Name,
		StartingBalance: entity.StartingBalance,
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
		Id:         uuid.New().String(),
		Title:      vo.Title,
		Comment:    vo.Comment,
		Date:       vo.Date,
		Amount:     vo.Amount,
		CategoryId: vo.CategoryId,
		AccountId:  vo.AccountId,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
}

func BookingEntityToVo(entity model.Booking) model.BookingVo {
	return model.BookingVo{
		Id:         entity.Id,
		Title:      entity.Title,
		Comment:    entity.Comment,
		Date:       entity.Date,
		Amount:     entity.Amount,
		CategoryId: entity.CategoryId,
		AccountId:  entity.AccountId,
		Created:    entity.Created,
		Updated:    entity.Updated,
	}
}

func BookingEntitiesToVos(entities []model.Booking) []model.BookingVo {
	vos := make([]model.BookingVo, 0, len(entities))
	for _, entity := range entities {
		vos = append(vos, BookingEntityToVo(entity))
	}
	return vos
}
