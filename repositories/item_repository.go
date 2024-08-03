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
	Delete(itemId uint) error
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

func (r *ItemMemoryRepository) Delete(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
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

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// FindAll implements IItemRepository.
func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (r *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, itemId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

func (r *ItemRepository) Delete(itemId uint) error {
	deleteItem, err := r.FindById(itemId)
	if err != nil {
		return err
	}

	result := r.db.Delete(&deleteItem) // 論理削除(gormデフォルト)で行いたいため、Unscoped().Delete(&deleteItem)とはしない
	if result.Error != nil {
		return result.Error
	}
	return nil
}
