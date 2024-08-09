package stockrepositorydomain

import stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"

type StockRepository interface {
	FindBySku(sku string) (*stockentity.Stock, error)
	Save(stock *stockentity.Stock) (*stockentity.Stock, error)
	GetAll() ([]*stockentity.Stock, error)
	Delete(sku string) error
}
