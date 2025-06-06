package repositories

import (
	"go-crud/config"
	"go-crud/models"
	"time"

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

func NewItemRepository() ItemRepository {
	return &itemRepository{db: config.DB}
}

func (r *itemRepository) CreateItem(item *models.Item) error {
	item.ID = uuid.New()
	item.BaseModel.CreatedAt = time.Now()
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
	item.BaseModel.UpdatedAt = time.Now()
	item.BaseModel.UpdatedBy = item.BaseModel.CreatedBy
	return r.db.Save(item).Error
}
