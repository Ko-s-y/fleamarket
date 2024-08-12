package main

import (
	"encoding/json"
	"fleamarket/infra"
	"fleamarket/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, UserID: 1},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト２", SoldOut: true, UserID: 1},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト３", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	// テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの実行結果の取得
	var res map[string][]models.Item
	err := json.Unmarshal([]byte(w.Body.String()), &res)
	if err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code, "Expected status code to be 200")
	assert.Equal(t, 3, len(res["data"]), "Expected data length to be 3")
}
