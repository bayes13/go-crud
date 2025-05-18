package helpers

import (
	"go-crud/models"
	"go-crud/models/dto"
)

func ToItemEntity(item *dto.ItemDto) *models.Item {
	return &models.Item{
		ID:          ParseUUID(item.ID),
		Code:        item.Code,
		Name:        item.Name,
		Description: item.Description,
		CreatedBy:   item.CreatedBy,
		Type:        item.Type,
		Category:    item.Category,
		UnitType:    item.UnitType,
		Enable:      item.Enable,
	}

}

func ToItemDto(item *models.Item) *dto.ItemDto {
	return &dto.ItemDto{
		ID:          item.ID.String(),
		Code:        item.Code,
		Name:        item.Name,
		Description: item.Description,
		CreatedBy:   item.CreatedBy,
		Type:        item.Type,
		Category:    item.Category,
		UnitType:    item.UnitType,
		Enable:      item.Enable,
		UpdatedBy:   item.UpdatedBy,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func ToItemDtoList(items *[]models.Item) []dto.ItemDto {
	dtoList := make([]dto.ItemDto, 0, len(*items))
	for _, item := range *items {
		dto := ToItemDto(&item)
		dtoList = append(dtoList, *dto)
	}
	return dtoList
}
