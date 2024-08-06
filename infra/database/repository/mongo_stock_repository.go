package mongostockrepository

import (
	"context"

	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	stockdocument "github.com/1EG/oms-inventory-go/infra/database/document"
	stockdatamapper "github.com/1EG/oms-inventory-go/infra/database/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	database *mongo.Database
}

func (repository *Repository) Save(stock *stockentity.Stock) (*stockentity.Stock, error) {
	collection := repository.database.Collection("stocks")

	filters := bson.D{{Key: "sku", Value: stock.Sku}}
	update := bson.D{{Key: "$set", Value: stockdatamapper.ToDocument(stock)}}
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filters, update, opts)

	if err != nil {
		return nil, err
	}

	return stock, nil
}

func (repository *Repository) GetAll() ([]*stockentity.Stock, error) {
	collection := repository.database.Collection("stocks")

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var documents []*stockdocument.Document

	if err = cursor.All(context.TODO(), &documents); err != nil {
		return nil, err
	}

	var stocks []*stockentity.Stock

	for _, document := range documents {
		stocks = append(stocks, stockdatamapper.ToDomain(document))
	}

	return stocks, nil
}

func Build(database *mongo.Database) *Repository {
	return &Repository{
		database: database,
	}
}
