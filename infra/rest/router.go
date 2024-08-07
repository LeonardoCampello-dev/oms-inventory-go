package router

import (
	stockhandler "github.com/1EG/oms-inventory-go/infra/rest/handler/stock"
	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	router := gin.Default()

	router.GET("/stock", stockhandler.GetAll)
	router.PUT("/stock/:sku", stockhandler.Save)
	router.DELETE("/stock/:sku", stockhandler.Delete)

	router.Run("localhost:8080")
}
