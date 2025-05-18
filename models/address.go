package models

import (
	"github.com/google/uuid"
)

type Address struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Label       string
	FullAddress string
	Contacts    []Contact `gorm:"foreignKey:AddressId;constrain:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
