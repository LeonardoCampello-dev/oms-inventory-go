package mongodatamapper

import (
	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	mongodocument "github.com/1EG/oms-inventory-go/infra/database/document"
)

func StockDocumentToDomain(document *mongodocument.Document) *stockentity.Stock {
	return &stockentity.Stock{
		Sku:       document.Sku,
		AccountId: document.AccountId,
		Quantity:  document.Quantity,
		Id:        document.Id,
	}
}

func StockToDocument(domain *stockentity.Stock) *mongodocument.Document {
	return &mongodocument.Document{
		Sku:       domain.Sku,
		AccountId: domain.AccountId,
		Quantity:  domain.Quantity,
		Id:        domain.Id,
	}
}
