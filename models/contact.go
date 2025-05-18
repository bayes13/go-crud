package models

import "github.com/google/uuid"

type Contact struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	FullName    string
	PhoneNumber string
	Type        string `gorm:"type:enum('INTERNAl','CUSTOMER','VENDOR')"`
	AddressID   uuid.UUID
}
