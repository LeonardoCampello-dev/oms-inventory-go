package getallstockshandler

import (
	"net/http"

	servicefactory "github.com/1EG/oms-inventory-go/infra/factory/service"
	"github.com/gin-gonic/gin"
)

func GetAll(context *gin.Context) {
	stockService := servicefactory.BuildStock()

	stocks, err := stockService.GetAll()

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.IndentedJSON(http.StatusOK, stocks)
}
