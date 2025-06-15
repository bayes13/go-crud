package dto

import "time"

type BaseModelDto struct {
	CreatedBy string    `json:"createdBy,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty"`
	CreatedAt time.Time `json:"createdDate"`
	UpdatedAt time.Time `json:"updatedDate"`
}
