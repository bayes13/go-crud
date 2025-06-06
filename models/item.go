package models

import (
	"github.com/google/uuid"
)

type Item struct {
	BaseModel

	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Code        string
	Name        string
	Description string
	Type        string `gorm:"type:text;check:type IN ('PRODUCT','SERVICE','COMBO','MATERIAL')"`
	Category    string
	UnitType    string
	Enable      bool
}
