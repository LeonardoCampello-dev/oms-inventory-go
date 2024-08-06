package repositoryfactory

import (
	"os"

	stockrepository "github.com/1EG/oms-inventory-go/domain/stock/repository"
	mongodb "github.com/1EG/oms-inventory-go/infra/database/mongo"
	mongostockrepository "github.com/1EG/oms-inventory-go/infra/database/repository"
	"github.com/joho/godotenv"
)

func BuildStock() stockrepository.RepositoryInterface {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	_, database, err := mongodb.Connect(os.Getenv("MONGO_URI"), "oms-inventory")

	if err != nil {
		panic(err)
	}

	return mongostockrepository.Build(database)
}
