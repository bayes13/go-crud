package models

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Code        string
	Name        string
	Description string
	Type        string `gorm:"type:enum('PRODUCT','SERVICE','COMBO','MATERIAL')"`
	Category    string
	UnitType    string
	Enable      bool
	CreatedBy   string
	UpdatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
