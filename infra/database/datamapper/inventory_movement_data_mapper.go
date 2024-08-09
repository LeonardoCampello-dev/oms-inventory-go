package mongodatamapper

import (
	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	stockvalueobject "github.com/1EG/oms-inventory-go/domain/stock/value_object"
	mongodocument "github.com/1EG/oms-inventory-go/infra/database/document"
)

func InventoryMovementDocumentToDomain(document *mongodocument.InventoryMovement) *stockentity.InventoryMovement {
	return &stockentity.InventoryMovement{
		AccountId:        document.AccountId,
		Sku:              document.Sku,
		Id:               document.Id,
		MovementType:     document.MovementType,
		PreviousQuantity: document.PreviousQuantity,
		CurrentQuantity:  document.CurrentQuantity,
		Order: stockvalueobject.InventoryMovementOrder{
			Id:         document.Order.Id,
			Sequential: document.Order.Sequential,
		},
	}
}

func InventoryMovementToDocument(domain *stockentity.InventoryMovement) *mongodocument.InventoryMovement {
	return &mongodocument.InventoryMovement{
		AccountId:        domain.AccountId,
		Sku:              domain.Sku,
		Id:               domain.Id,
		MovementType:     domain.MovementType,
		PreviousQuantity: domain.PreviousQuantity,
		CurrentQuantity:  domain.CurrentQuantity,
		Order: mongodocument.InventoryMovementOrder{
			Id:         domain.Order.Id,
			Sequential: domain.Order.Sequential,
		},
	}
}
