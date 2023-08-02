package http

import (
	"github.com/gin-gonic/gin"
)

func (d *Delivery) initRouter() *gin.Engine {

	var router = gin.New()

	d.routerInventory(router.Group("/inventory"))

	return router
}

func (d *Delivery) routerInventory(router *gin.RouterGroup) {
	//router.POST("/", d.CreateInventory)
	//router.PUT("/:id", d.UpdateInventory)
	//router.DELETE("/:id", d.DeleteInventory)
	//router.GET("/", d.ListInventory)
	//router.GET("/:id", d.ReadInventoryByID)
}
