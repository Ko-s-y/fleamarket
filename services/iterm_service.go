package services

import (
	"fleamarket/models"
	"fleamarket/repositories"
)

// IItemService インターフェースは、アイテムサービスの契約を定義
type IItemService interface {
	FindAll() (*[]models.Item, error) // すべてのアイテムを取得するmethod
}

// ItemService 構造体は、IItemService を実装
type ItemService struct {
	repository repositories.IItemRepository // アイテムリポジトリを保持
}

// NewItemService 関数は、ItemService のインスタンスを初期化
func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository} // 新しい ItemService インスタンスを返す
}

// FindAll メソッドは、リポジトリの FindAll() メソッドを呼び出し、その結果を返す
func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}
