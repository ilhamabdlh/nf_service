package service

import (
	"nft_service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/items", handlers.GetItems)
		api.GET("/items/:id", handlers.GetItemByID)
		api.POST("/items", handlers.CreateItem)
		api.PUT("/items/:id", handlers.UpdateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)
		api.POST("/purchase/:id", handlers.PurchaseItem)
	}
}
