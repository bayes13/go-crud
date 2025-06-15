package dto

type LocationDto struct {
	BaseModelDto
	ID       string     `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Category string     `json:"category,omitempty"`
	Type     string     `json:"type,omitempty"`
	Enable   bool       `json:"enable,omitempty"`
	Address  AddressDto `json:"address"`
}
