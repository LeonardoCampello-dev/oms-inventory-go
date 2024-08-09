package stockentity

import (
	stockvalueobject "github.com/1EG/oms-inventory-go/domain/stock/value_object"
	"github.com/google/uuid"
)

type InventoryMovement struct {
	AccountId        string `json:"-"`
	Sku              string `json:"sku"`
	Id               string `json:"id"`
	MovementType     string `json:"movementType"`
	PreviousQuantity int    `json:"previousQuantity"`
	CurrentQuantity  int    `json:"currentQuantity"`
	Order            stockvalueobject.InventoryMovementOrder
}

func BuildManualInventoryMovement(accountId string, sku string, previousQuantity int, currentQuantity int) *InventoryMovement {
	return &InventoryMovement{
		AccountId:        accountId,
		MovementType:     "MANUAL",
		Sku:              sku,
		PreviousQuantity: previousQuantity,
		CurrentQuantity:  currentQuantity,
		Id:               uuid.NewString(),
	}
}

func BuildOrderInventoryMovement(accountId string, sku string, previousQuantity int, currentQuantity int, order stockvalueobject.InventoryMovementOrder) *InventoryMovement {
	return &InventoryMovement{
		AccountId:        accountId,
		MovementType:     "ORDER",
		Sku:              sku,
		PreviousQuantity: previousQuantity,
		CurrentQuantity:  currentQuantity,
		Order:            order,
		Id:               uuid.NewString(),
	}
}
