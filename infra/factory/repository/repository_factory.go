package repositoryfactory

import (
	"os"

	stockrepositorydomain "github.com/1EG/oms-inventory-go/domain/stock/repository"
	"github.com/1EG/oms-inventory-go/infra/database/mongodatabase"
	mongorepository "github.com/1EG/oms-inventory-go/infra/database/repository"
	"github.com/joho/godotenv"
)

func BuildStock() stockrepositorydomain.StockRepository {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	_, database, err := mongodatabase.Connect(os.Getenv("MONGO_URI"), "oms-inventory")

	if err != nil {
		panic(err)
	}

	return mongorepository.BuildStockRepository(database)
}

func BuildInventoryMovement() stockrepositorydomain.InventoryMovementRepository {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	_, database, err := mongodatabase.Connect(os.Getenv("MONGO_URI"), "oms-inventory")

	if err != nil {
		panic(err)
	}

	return mongorepository.BuildInventoryMovementRepository(database)
}
