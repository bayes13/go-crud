package models

import (
	"github.com/google/uuid"
)

type Contact struct {
	BaseModel

	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	FullName    string
	PhoneNumber string
	Type        string `gorm:"type:text;check:type IN('INTERNAl','CUSTOMER','VENDOR')"`
	AddressID   *uuid.UUID
	Address     *Address `gorm:"foreign_key:AddressID"`
}
