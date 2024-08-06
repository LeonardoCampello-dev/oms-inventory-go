package deletestockhandler

import (
	"net/http"

	servicefactory "github.com/1EG/oms-inventory-go/infra/factory/service"
	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	sku := context.Param("sku")

	stockService := servicefactory.BuildStock()

	err := stockService.Delete(sku)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.Status(http.StatusNoContent)
}
