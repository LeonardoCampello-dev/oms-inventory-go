package stockhandler

import (
	"net/http"

	servicefactory "github.com/1EG/oms-inventory-go/infra/factory/service"
	"github.com/gin-gonic/gin"
)

type RequestDTO struct {
	Quantity int `json:"quantity" binding:"required"`
}

func Save(context *gin.Context) {
	sku := context.Param("sku")

	stockService := servicefactory.BuildStock()

	var requestDTO RequestDTO

	if err := context.ShouldBindJSON(&requestDTO); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	stock, err := stockService.Save(sku, requestDTO.Quantity, "AccountId")

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"sku": stock.Sku, "quantity": stock.Quantity})
}
