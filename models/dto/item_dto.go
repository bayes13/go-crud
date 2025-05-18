package dto

import "time"

type ItemDto struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	UnitType    string    `json:"unitType"`
	Enable      bool      `json:"enable"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"`
	CreatedAt   time.Time `json:"createdDate"`
	UpdatedAt   time.Time `json:"updatedDate"`
}
