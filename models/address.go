package models

import (
	"github.com/google/uuid"
)

type Address struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Label       string    `json:"label"`
	FullAddress string    `json:"author"`
	Contacts    []Contact `gorm:"foreignKey:AddressId;constrain:OnUpdate:CASCADE,OnDelete:SET NULL" json:"contacts"`
}
