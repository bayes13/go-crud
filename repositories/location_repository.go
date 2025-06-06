package repositories

import (
	"go-crud/config"
	"go-crud/models"

	"gorm.io/gorm"
)

type LocationRepository interface {
	CreateLocation(location *models.Location) error
	FindLocation(location *models.Location) ([]models.Location, error)
	UpdateLocation(location *models.Location) error
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{db: config.DB}

}

func (r *locationRepository) CreateLocation(location *models.Location) error {
	return nil
}

func (r *locationRepository) FindLocation(location *models.Location) ([]models.Location, error) {
	var locations []models.Location
	return locations, nil
}

func (r *locationRepository) UpdateLocation(location *models.Location) error {
	return nil
}
