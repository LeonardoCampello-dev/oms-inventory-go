package stockservice

import (
	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	stockrepository "github.com/1EG/oms-inventory-go/domain/stock/repository"
)

type Service struct {
	stockRepository stockrepository.RepositoryInterface
}

func Build(stockRepository stockrepository.RepositoryInterface) *Service {
	return &Service{stockRepository: stockRepository}
}

func (stockService *Service) GetAll() ([]*stockentity.Stock, error) {
	stocks, err := stockService.stockRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (stockService *Service) Save(sku string, quantity int, accountId string) (*stockentity.Stock, error) {
	stock := stockentity.Build(accountId, sku, quantity)

	createdStock, err := stockService.stockRepository.Save(stock)

	if err != nil {
		return nil, err
	}

	return createdStock, nil
}
