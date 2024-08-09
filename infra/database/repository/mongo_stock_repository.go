package mongorepository

import (
	"context"

	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	mongodatamapper "github.com/1EG/oms-inventory-go/infra/database/datamapper"
	mongodocument "github.com/1EG/oms-inventory-go/infra/database/document"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StockRepository struct {
	database *mongo.Database
}

func (repository *StockRepository) FindBySku(sku string) (*stockentity.Stock, error) {
	collection := repository.database.Collection("stocks")

	var result mongodocument.Document

	err := collection.FindOne(context.TODO(), bson.D{{Key: "sku", Value: sku}}).Decode(&result)

	if err != nil {
		return nil, err
	}

	return mongodatamapper.StockDocumentToDomain(&result), nil
}

func (repository *StockRepository) Delete(sku string) error {
	collection := repository.database.Collection("stocks")

	_, err := collection.DeleteOne(context.TODO(), bson.D{{Key: "sku", Value: sku}})

	if err != nil {
		return err
	}

	return nil
}

func (repository *StockRepository) Save(stock *stockentity.Stock) (*stockentity.Stock, error) {
	collection := repository.database.Collection("stocks")

	filters := bson.D{{Key: "sku", Value: stock.Sku}}
	update := bson.D{{Key: "$set", Value: mongodatamapper.StockToDocument(stock)}}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filters, update, opts)

	if err != nil {
		return nil, err
	}

	return stock, nil
}

func (repository *StockRepository) GetAll() ([]*stockentity.Stock, error) {
	collection := repository.database.Collection("stocks")

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var documents []*mongodocument.Document

	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}

	var stocks []*stockentity.Stock

	for _, document := range documents {
		stocks = append(stocks, mongodatamapper.StockDocumentToDomain(document))
	}

	return stocks, nil
}

func BuildStockRepository(database *mongo.Database) *StockRepository {
	return &StockRepository{
		database: database,
	}
}
