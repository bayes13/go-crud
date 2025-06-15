package dto

type ItemDto struct {
	BaseModelDto
	ID          string `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Category    string `json:"category,omitempty"`
	UnitType    string `json:"unitType,omitempty"`
	Enable      bool   `json:"enable,omitempty"`
}
