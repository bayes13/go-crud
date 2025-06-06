package models

import (
	"github.com/google/uuid"
)

type Location struct {
	BaseModel

	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string
	Category  string
	Type      string `gorm:"type:text;check:type IN('INTERNAL','SCRAP','ADJUSTMENT','SUPLIER_ADD','CUSTOMER_ADD')"`
	Enabled   bool
	Deleted   bool
	Address   Address `gorm:"foreign_key:AddressID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AddressID uuid.UUID
}
