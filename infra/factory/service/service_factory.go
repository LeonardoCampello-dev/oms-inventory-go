package servicefactory

import (
	stockservice "github.com/1EG/oms-inventory-go/domain/stock/service"
	repositoryfactory "github.com/1EG/oms-inventory-go/infra/factory/repository"
)

func BuildStock() *stockservice.Service {
	return stockservice.Build(repositoryfactory.BuildStock())
}
