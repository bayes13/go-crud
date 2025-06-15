package repositories

import (
	"go-crud/config"
	"go-crud/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationRepository interface {
	CreateLocation(location *models.Location) error
	FindLocation(location *models.Location) ([]models.Location, error)
	UpdateLocation(location *models.Location) error
	FindLocationById(id uuid.UUID) (models.Location, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{db: config.DB}

}

func (r *locationRepository) CreateLocation(location *models.Location) error {
	location.ID = uuid.New()
	location.BaseModel.CreatedAt = time.Now()
	return r.db.Create(location).Error
}

func (r *locationRepository) FindLocation(location *models.Location) ([]models.Location, error) {
	var locations []models.Location
	query := r.db.Table("locations").
		Select("*").
		Joins("JOIN addresses addr on addr.id=locations.address_id").
		Joins("JOIN contacts con on con.address_id=addresses.id").
		Scan(&locations)

	err := query.Scan(&locations).Error

	return locations, err
}

func (r *locationRepository) UpdateLocation(location *models.Location) error {
	location.BaseModel.UpdatedAt = time.Now()
	location.BaseModel.UpdatedBy = location.BaseModel.CreatedBy
	return r.db.Save(location).Error
}

func (r *locationRepository) FindLocationById(id uuid.UUID) (models.Location, error) {
	var location models.Location
	err := r.db.First(&location, id).Error

	return location, err
}
