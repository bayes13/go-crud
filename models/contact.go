package models

import "github.com/google/uuid"

type Contact struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	FullName    string    `json:"fullName"`
	PhoneNumber string    `json:"phoneNumber"`
	Type        string    `gorm:"type:enum('INTERNAl','CUSTOMER','VENDOR')" json:"type"`
	AddressID   uuid.UUID `json:"addressId"`
}
