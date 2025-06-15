package dto

type ContactDto struct {
	BaseModelDto
	ID          string `json:"id,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Type        string `json:"type,omitempty"`
}
