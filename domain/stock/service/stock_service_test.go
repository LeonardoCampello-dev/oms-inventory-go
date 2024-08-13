package stockservice_test

import (
	"testing"

	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	stockservice "github.com/1EG/oms-inventory-go/domain/stock/service"
	mock_repository "github.com/1EG/oms-inventory-go/test/mock/repository"
	"github.com/bxcodec/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	gomockController := gomock.NewController(t)

	defer gomockController.Finish()

	stockRepository := mock_repository.NewMockStockRepository(gomockController)
	movementRepository := mock_repository.NewMockInventoryMovementRepository(gomockController)

	stockService := stockservice.Build(stockRepository, movementRepository)

	t.Run("should save new stock and create inventory movement", func(t *testing.T) {
		sku := faker.UUIDHyphenated()
		accountId := faker.UUIDHyphenated()
		quantity := 10

		stockRepository.EXPECT().FindBySku(sku).Return(nil, nil)
		stockRepository.EXPECT().Save(gomock.Any()).Return(&stockentity.Stock{Sku: sku, Quantity: quantity, AccountId: accountId}, nil)
		movementRepository.EXPECT().Save(gomock.Any()).Return(nil, nil)

		result, err := stockService.Save(sku, quantity, accountId)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, sku, result.Sku)
		assert.Equal(t, quantity, result.Quantity)
		assert.Equal(t, accountId, result.AccountId)
	})
}
