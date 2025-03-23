package repositories

import (
	"errors"
	"fleamarket/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{
		items: items,
	}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(id uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("item not found")
}
