package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string, dbName string) (*mongo.Client, *mongo.Database, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, nil, err
	}

	fmt.Println("You successfully connected to MongoDB!")

	db := client.Database(dbName)

	return client, db, nil
}
