package dto

type AddressDto struct {
	BaseModelDto
	ID          string       `json:"id,omitempty"`
	Label       string       `json:"label,omitempty"`
	FullAddress string       `json:"fullAddress,omitempty"`
	Contacts    []ContactDto `json:"contacts"`
}
