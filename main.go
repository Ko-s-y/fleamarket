package main

import (
	"fleamarket/controllers"
	"fleamarket/infra"
	"fleamarket/models"
	"fleamarket/repositories"
	"fleamarket/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	log.Println(os.Getenv("ENV"))

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
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
