package mapper

import (
	"github.com/koloo91/model"
)

func AccountVoToEntity(vo model.AccountVo) model.Account {
	return model.Account{
		Id:      vo.Id,
		Name:    vo.Name,
		Created: vo.Created,
		Updated: vo.Updated,
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
