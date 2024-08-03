package main

import (
	"fleamarket/infra"
	"fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("Failed to mograte database")
	}
}
