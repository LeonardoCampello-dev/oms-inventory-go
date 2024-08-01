package router

import (
	getallstockshandler "github.com/1EG/oms-inventory-go/infra/rest/handler"
	"github.com/gin-gonic/gin"
)

func Bootstrap() {
	router := gin.Default()

	router.GET("/stock", getallstockshandler.GetAll)

	router.Run("localhost:8080")
}
