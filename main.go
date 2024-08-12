package main

import (
	"fleamarket/controllers"
	"fleamarket/infra"
	"fleamarket/middlewares"
	"fleamarket/repositories"
	"fleamarket/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// item
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	// itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	// auth
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthServiceController(authService)
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	// authmiddlewareを利用した定義
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	itemRouterWithAuth.POST("", itemController.Create)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)

	r.Run("localhost:8080")
}
