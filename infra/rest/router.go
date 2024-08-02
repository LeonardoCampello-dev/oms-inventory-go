package router

import (
	getallstockshandler "github.com/1EG/oms-inventory-go/infra/rest/handler/get_all_stocks_handler"
	savestockhandler "github.com/1EG/oms-inventory-go/infra/rest/handler/save_stock_handler"
	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	router := gin.Default()

	router.GET("/stock", getallstockshandler.GetAll)
	router.PUT("/stock/:sku", savestockhandler.Save)

	router.Run("localhost:8080")
}
