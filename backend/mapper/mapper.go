package mapper

import (
	"github.com/google/uuid"
	"github.com/koloo91/model"
	"time"
)

func AccountVoToEntity(vo model.AccountVo) model.Account {
	return model.Account{
		Id:      uuid.New().String(),
		Name:    vo.Name,
		Created: time.Now(),
		Updated: time.Now(),
	}
}

func AccountEntityToVo(entity model.Account) model.AccountVo {
	return model.AccountVo{
		Id:      entity.Id,
		Name:    entity.Name,
		Created: entity.Created,
		Updated: entity.Updated,
	}
}

func AccountEntitiesToVos(entities []model.Account) []model.AccountVo {
	vos := make([]model.AccountVo, len(entities))
	for _, entity := range entities {
		vos = append(vos, AccountEntityToVo(entity))
	}
	return vos
}
