package models

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `gorm:"type:enum('PRODUCT','SERVICE','COMBO','MATERIAL')" json:"type"`
	Category    string    `json:"category"`
	UnitType    string    `gorm:"type:enum('PCS','LITER','METER')" json:"unitType"`
	Enable      bool      `json:"enable"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"`
	CreatedAt   time.Time
	UpdatedTime time.Time
}
