package models

import (
	"github.com/google/uuid"
)

type Location struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string
	Category  string
	Type      string `gorm:"type:enum('INTERNAL','SCRAP','ADJUSTMENT','SUPLIER_ADD','CUSTOMER_ADD')"`
	Enabled   bool
	Deleted   bool
	Address   Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AddressID uuid.UUID
}
