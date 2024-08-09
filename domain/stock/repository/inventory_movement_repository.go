package stockrepositorydomain

import stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"

type InventoryMovementRepository interface {
	Save(inventoryMovement *stockentity.InventoryMovement) (*stockentity.InventoryMovement, error)
}
