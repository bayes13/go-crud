package dto

import "time"

type ItemDto struct {
	ID          string    `json:"id,omitempty"`
	Code        string    `json:"code,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type,omitempty"`
	Category    string    `json:"category,omitempty"`
	UnitType    string    `json:"unitType,omitempty"`
	Enable      bool      `json:"enable,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty"`
	UpdatedBy   string    `json:"updatedBy,omitempty"`
	CreatedAt   time.Time `json:"createdDate,omitempty"`
	UpdatedAt   time.Time `json:"updatedDate,omitempty"`
}
