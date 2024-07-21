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
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 1000, Description: "説明2", SoldOut: false},
		{ID: 3, Name: "商品3", Price: 1000, Description: "説明3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	IItemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(IItemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
