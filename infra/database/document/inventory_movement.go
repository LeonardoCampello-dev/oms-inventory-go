package mongodocument

type InventoryMovementOrder struct {
	Id         string `bson:"id"`
	Sequential string `bson:"sequential"`
}

type InventoryMovement struct {
	AccountId        string `bson:"accountId"`
	Sku              string `bson:"sku"`
	Id               string `bson:"id"`
	MovementType     string `bson:"movementType"`
	PreviousQuantity int    `bson:"previousQuantity"`
	CurrentQuantity  int    `bson:"currentQuantity"`
	Order            InventoryMovementOrder
}
