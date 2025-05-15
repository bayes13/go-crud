package repositories

import (
	"go-crud/config"
	"go-crud/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *models.Item) error
	FindItem(item *models.Item) ([]models.Item, error)
	UpdateItem(item *models.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

func newItemRepository() ItemRepository {
	return &itemRepository{db: config.DB}
}

func (r *itemRepository) CreateItem(item *models.Item) error {
	item.ID = uuid.New()
	return r.db.Create(item).Error
}

func (r *itemRepository) FindItem(item *models.Item) ([]models.Item, error) {
	var items []models.Item
	query := r.db.Model(item)
	if err := query.Where(item).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *itemRepository) UpdateItem(item *models.Item) error {
	return r.db.Save(item).Error
}
