package models

import (
	"github.com/google/uuid"
)

type Location struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Type      string    `gorm:"type:enum('INTERNAL','SCRAP','ADJUSTMENT','SUPLIER_ADD','CUSTOMER_ADD')" json:"type"`
	Enabled   bool      `json:"enabled"`
	Deleted   bool      `json:"deleted"`
	Address   Address   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AddressID uuid.UUID `json:"addressid"`
}
