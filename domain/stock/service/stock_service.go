package stockservice

import (
	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	stockrepositorydomain "github.com/1EG/oms-inventory-go/domain/stock/repository"
)

type StockService struct {
	inventoryMovementRepository stockrepositorydomain.InventoryMovementRepository
	stockRepository             stockrepositorydomain.StockRepository
}

func Build(stockRepository stockrepositorydomain.StockRepository, inventoryMovementRepository stockrepositorydomain.InventoryMovementRepository) *StockService {
	return &StockService{stockRepository: stockRepository, inventoryMovementRepository: inventoryMovementRepository}
}

func (service *StockService) Delete(sku string) error {
	err := service.stockRepository.Delete(sku)

	if err != nil {
		return err
	}

	return nil
}

func (service *StockService) GetAll() ([]*stockentity.Stock, error) {
	stocks, err := service.stockRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (service *StockService) Save(sku string, quantity int, accountId string) (*stockentity.Stock, error) {
	existingStock, err := service.stockRepository.FindBySku(sku)

	if err != nil {
		return nil, err
	}

	stock := stockentity.Build(accountId, sku, quantity)

	savedStock, err := service.stockRepository.Save(stock)

	if err != nil {
		return nil, err
	}

	previousStockQuantity := 0

	if existingStock != nil {
		previousStockQuantity = existingStock.Quantity
	}

	inventoryMovement := stockentity.BuildManualInventoryMovement(accountId, sku, previousStockQuantity, quantity)

	if _, err := service.inventoryMovementRepository.Save(inventoryMovement); err != nil {
		return nil, err
	}

	return savedStock, nil
}
