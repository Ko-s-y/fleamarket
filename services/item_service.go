package services

import (
	"fleamarket/dto"
	"fleamarket/models"
	"fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(id uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{
		repository: repository,
	}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(id uint) (*models.Item, error) {
	return s.repository.FindById(id)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	item := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}

	return s.repository.Create(item)
}

func (s *ItemService) Update(id uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}

	return s.repository.Update(*targetItem)
}
