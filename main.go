package main

import (
	"fleamarket/controllers"
	"fleamarket/models"
	"fleamarket/repositories"
	"fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "Item 1", Price: 1000, Description: "Description 1", SoldOut: false},
		{ID: 2, Name: "Item 2", Price: 2000, Description: "Description 2", SoldOut: true},
		{ID: 3, Name: "Item 3", Price: 3000, Description: "Description 3", SoldOut: false},
	}
	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
