package services

import (
	"go-crud/helpers"
	"go-crud/models"
	"go-crud/repositories"
)

type LocationService interface {
	CreateLocation(location *models.Location) bool
	FindLocation(location *models.Location) []models.Location
	UpdateLocation(location *models.Location) bool
	DisableLocation(id string, isEnable bool) bool
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (l *locationService) CreateLocation(location *models.Location) bool {
	if err := l.repo.CreateLocation(location); err != nil {
		return false
	}
	return true
}

func (l *locationService) DisableLocation(id string, isEnable bool) bool {
	locationId := helpers.ParseUUID(id)
	if location, err := l.repo.FindLocationById(locationId); err != nil {
		return false
	} else {
		location.Enable = isEnable
		return l.UpdateLocation(&location)
	}
}

func (l *locationService) FindLocation(location *models.Location) []models.Location {
	if locations, err := l.repo.FindLocation(location); err != nil {
		return []models.Location{}
	} else {
		return locations
	}
}

func (l *locationService) UpdateLocation(location *models.Location) bool {
	if err := l.repo.UpdateLocation(location); err != nil {
		return false
	}
	return true
}
