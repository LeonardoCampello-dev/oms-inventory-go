package stockdatamapper

import (
	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	stockdocument "github.com/1EG/oms-inventory-go/infra/database/document"
)

func ToDomain(document *stockdocument.Document) *stockentity.Stock {
	return &stockentity.Stock{
		Sku:       document.Sku,
		AccountId: document.AccountId,
		Quantity:  document.Quantity,
		Id:        document.Id,
	}
}

func ToDocument(domain *stockentity.Stock) *stockdocument.Document {
	return &stockdocument.Document{
		Sku:       domain.Sku,
		AccountId: domain.AccountId,
		Quantity:  domain.Quantity,
		Id:        domain.Id,
	}
}
