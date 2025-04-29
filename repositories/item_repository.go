package repositories

import (
	"errors"
	"fleamarket/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(item models.Item) (*models.Item, error)
	Update(item models.Item) (*models.Item, error)
	Delete(id uint) error
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

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(item models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == item.ID {
			r.items[i] = item
			return &r.items[i], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (r *ItemMemoryRepository) Delete(id uint) error {
	for i, v := range r.items {
		if v.ID == id {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
