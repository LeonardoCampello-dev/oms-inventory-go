package stockentity

import "github.com/google/uuid"

type Stock struct {
	Sku       string `json:"sku"`
	AccountId string `json:"-"`
	Quantity  int    `json:"quantity"`
	Id        string `json:"id"`
}

type StockOption func(*Stock)

func WithID(id string) StockOption {
	return func(stock *Stock) {
		stock.Id = id
	}
}

func Build(accountId string, sku string, quantity int, options ...StockOption) *Stock {
	stock := &Stock{
		Sku:       sku,
		AccountId: accountId,
		Quantity:  quantity,
		Id:        uuid.NewString(),
	}

	for _, option := range options {
		option(stock)
	}

	return stock
}
