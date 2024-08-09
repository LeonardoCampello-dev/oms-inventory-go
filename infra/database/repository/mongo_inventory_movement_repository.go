package mongorepository

import (
	"context"

	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	mongodatamapper "github.com/1EG/oms-inventory-go/infra/database/datamapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInventoryMovementRepository struct {
	database *mongo.Database
}

func (repository *MongoInventoryMovementRepository) Save(inventoryMovement *stockentity.InventoryMovement) (*stockentity.InventoryMovement, error) {
	collection := repository.database.Collection("inventoryMovement")

	filters := bson.D{{Key: "sku", Value: inventoryMovement.Sku}}
	update := bson.D{{Key: "$set", Value: mongodatamapper.InventoryMovementToDocument(inventoryMovement)}}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filters, update, opts)

	if err != nil {
		return nil, err
	}

	return inventoryMovement, nil
}

func BuildInventoryMovementRepository(database *mongo.Database) *MongoInventoryMovementRepository {
	return &MongoInventoryMovementRepository{
		database: database,
	}
}
