package models

import (
	"github.com/google/uuid"
)

type Address struct {
	BaseModel

	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Label       string
	FullAddress string
	Contacts    []Contact `gorm:"foreign_key:AddressID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
