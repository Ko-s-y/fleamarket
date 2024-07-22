package repositories

import (
	"errors"
	"fleamarket/models"
)

// IItemRepository インターフェースは、アイテムリポジトリの契約を定義
type IItemRepository interface {
	FindAll() (*[]models.Item, error) // すべてのアイテムを取得するmethod
	FindById(itemId uint) (*models.Item, error)
}

// ItemMemoryRepository 構造体は、IItemRepository のメモリ内実装
type ItemMemoryRepository struct {
	items []models.Item // アイテムを保存するスライス
}

// NewItemMemoryRepository 関数は、指定されたアイテムのスライスで ItemMemoryRepository を初期化
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items} // 新しい ItemMemoryRepository インスタンスを返す
}

// FindAll メソッドは、リポジトリに保存されているすべてのアイテムを取得
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil // アイテムと nil エラーを返す
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}
