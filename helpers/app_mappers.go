package helpers

import (
	"go-crud/models"
	"go-crud/models/dto"
)

func ToItemEntity(item *dto.ItemDto) *models.Item {
	return &models.Item{
		BaseModel: models.BaseModel{
			CreatedBy: item.CreatedBy,
		},
		ID:          ParseUUID(item.ID),
		Code:        item.Code,
		Name:        item.Name,
		Description: item.Description,
		Type:        item.Type,
		Category:    item.Category,
		UnitType:    item.UnitType,
		Enable:      item.Enable,
	}

}

func ToItemDto(item *models.Item) *dto.ItemDto {
	return &dto.ItemDto{
		BaseModelDto: dto.BaseModelDto{
			CreatedBy: item.CreatedBy,
			UpdatedBy: item.BaseModel.UpdatedBy,
			CreatedAt: item.BaseModel.CreatedAt,
			UpdatedAt: item.BaseModel.UpdatedAt,
		},
		ID:          item.ID.String(),
		Code:        item.Code,
		Name:        item.Name,
		Description: item.Description,
		Type:        item.Type,
		Category:    item.Category,
		UnitType:    item.UnitType,
		Enable:      item.Enable,
	}
}

func toContactDto(contact *models.Contact) *dto.ContactDto {
	return &dto.ContactDto{
		BaseModelDto: dto.BaseModelDto{
			CreatedBy: contact.CreatedBy,
			UpdatedBy: contact.BaseModel.UpdatedBy,
			CreatedAt: contact.BaseModel.CreatedAt,
			UpdatedAt: contact.BaseModel.UpdatedAt,
		},
		ID:          contact.ID.String(),
		FullName:    contact.FullName,
		PhoneNumber: contact.PhoneNumber,
		Type:        contact.Type,
	}

}

func toAddresstDto(address *models.Address) *dto.AddressDto {
	return &dto.AddressDto{
		BaseModelDto: dto.BaseModelDto{
			CreatedBy: address.CreatedBy,
			UpdatedBy: address.BaseModel.UpdatedBy,
			CreatedAt: address.BaseModel.CreatedAt,
			UpdatedAt: address.BaseModel.UpdatedAt,
		},
		ID:          address.ID.String(),
		Label:       address.Label,
		FullAddress: address.FullAddress,
		Contacts:    ToContactDtoList(&address.Contacts),
	}

}
func toLocationDto(location *models.Location) *dto.LocationDto {
	return &dto.LocationDto{
		BaseModelDto: dto.BaseModelDto{
			CreatedBy: location.CreatedBy,
			UpdatedBy: location.BaseModel.UpdatedBy,
			CreatedAt: location.BaseModel.CreatedAt,
			UpdatedAt: location.BaseModel.UpdatedAt,
		},
		ID:       location.ID.String(),
		Name:     location.Name,
		Category: location.Category,
		Type:     location.Type,
		Enable:   location.Enable,
		Address:  *toAddresstDto(location.Address),
	}

}
func ToContactDtoList(contacts *[]models.Contact) []dto.ContactDto {
	dtoList := make([]dto.ContactDto, 0, len(*contacts))
	for _, contact := range *contacts {
		dto := toContactDto(&contact)
		dtoList = append(dtoList, *dto)
	}
	return dtoList
}

func ToItemDtoList(items *[]models.Item) []dto.ItemDto {
	dtoList := make([]dto.ItemDto, 0, len(*items))
	for _, item := range *items {
		dto := ToItemDto(&item)
		dtoList = append(dtoList, *dto)
	}
	return dtoList
}
