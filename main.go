package main

import (
	"fleamarket/controllers"
	"fleamarket/infra"
	"fleamarket/repositories"
	"fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := gin.Default()

	// item
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	// auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthServiceController(authService)
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)

	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
