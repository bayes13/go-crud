package services

import (
	"go-crud/helpers"
	"go-crud/models"
	"go-crud/repositories"
)

type ItemService interface {
	CreateItem(item *models.Item) bool
	FindItem(item *models.Item) []models.Item
	UpdateItem(item *models.Item) bool
	DisableItem(id string, isEnable bool) bool
}

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) CreateItem(item *models.Item) bool {
	if err := s.repo.CreateItem(item); err != nil {
		return false
	}
	return true
}

func (s *itemService) FindItem(item *models.Item) []models.Item {
	if val, err := s.repo.FindItem(item); err != nil {
		return []models.Item{}
	} else {
		return val
	}
}

func (s *itemService) UpdateItem(item *models.Item) bool {
	if err := s.repo.UpdateItem(item); err != nil {
		return false
	}
	return true
}

func (s *itemService) DisableItem(id string, isEnable bool) bool {
	var item *models.Item
	item.ID = helpers.ParseUUID(id)
	items := s.FindItem(item)
	if len(items) == 1 {
		items[0].Enable = isEnable
		return s.UpdateItem(item)
	} else {
		return false
	}
}
