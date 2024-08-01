package stockrepository

import stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"

type RepositoryInterface interface {
	Save(stock *stockentity.Stock) (*stockentity.Stock, error)
	GetAll() ([]*stockentity.Stock, error)
}
