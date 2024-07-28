package services

import (
	"fleamarket/dto"
	"fleamarket/models"
	"fleamarket/repositories"
)

// IItemService インターフェースは、アイテムサービスの契約を定義
type IItemService interface {
	FindAll() (*[]models.Item, error) // すべてのアイテムを取得するmethod
	FindById(itemId uint) (*models.Item, error)
	Create(CreateItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint) error
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

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.FindById(itemId)
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

func (s *ItemService) Delete(itemId uint) error {
	return s.repository.Delete(itemId)
}
