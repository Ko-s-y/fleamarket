package repositories

import (
	"errors"
	"fleamarket/models"

	"gorm.io/gorm"
)

// IItemRepository インターフェースは、アイテムリポジトリの契約を定義
type IItemRepository interface {
	FindAll() (*[]models.Item, error) // すべてのアイテムを取得するメソッド
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemID uint) error
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

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (r *ItemMemoryRepository) Delete(itemID uint) error {
	for i, v := range r.items {
		if v.ID == itemID {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}

// ItemRepository 構造体は、IItemRepository のデータベース実装
type ItemRepository struct {
	db *gorm.DB
}

// NewItemRepository 関数は、ItemRepository のインスタンスを初期化
func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

// Create implements IItemRepository.
func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// Delete implements IItemRepository.
func (r *ItemRepository) Delete(itemID uint) error {
	panic("unimplemented")
}

// FindAll implements IItemRepository.
func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	panic("unimplemented")
}

// Update implements IItemRepository.
func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	panic("unimplemented")
}

// FindById implements IItemRepository.
func (r *ItemRepository) FindById(itemID uint) (*models.Item, error) {
	panic("unimplemented")
}
