package repositoryfactory

import (
	stockrepository "github.com/1EG/oms-inventory-go/domain/stock/repository"
	mongodb "github.com/1EG/oms-inventory-go/infra/database/mongo"
	mongostockrepository "github.com/1EG/oms-inventory-go/infra/database/repository"
)

func BuildStock() stockrepository.RepositoryInterface {
	_, database, err := mongodb.Connect("TODO", "oms-inventory")

	if err != nil {
		panic(err)
	}

	return mongostockrepository.Build(database)
}
