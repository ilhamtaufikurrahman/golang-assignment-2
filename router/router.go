package router

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.DELETE("/orders/:OrderID", controllers.DeleteOrder)
	router.PUT("/orders/:OrderID", controllers.UpdateOrder)

	return router
}
