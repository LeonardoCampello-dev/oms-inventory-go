package getallstockshandler

import (
	"net/http"

	stockentity "github.com/1EG/oms-inventory-go/domain/stock/entity"
	"github.com/gin-gonic/gin"
)

func GetAll(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, []stockentity.Stock{{Sku: "123", AccountId: "123", Quantity: 1, Id: "123"}})
}
