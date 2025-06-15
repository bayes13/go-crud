package controlers

import (
	"go-crud/helpers"
	"go-crud/models/dto"
	"go-crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerApp struct {
	ItemService     services.ItemService
	LocationService services.LocationService
}

func NewControllerApp(itemService services.ItemService, locationService services.LocationService) *ControllerApp {
	return &ControllerApp{ItemService: itemService, LocationService: locationService}
}

func (c *ControllerApp) CreateItem(ctx *gin.Context) {
	var dto dto.ItemDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success := c.ItemService.CreateItem(helpers.ToItemEntity(&dto))
	if !success {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Create Item"})
		return
	}
	ctx.JSON(http.StatusCreated, dto)
}

func (c *ControllerApp) UpdateItem(ctx *gin.Context) {
	var dto dto.ItemDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success := c.ItemService.UpdateItem(helpers.ToItemEntity(&dto))
	if !success {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update Item"})
		return
	}
	ctx.JSON(http.StatusCreated, dto)

}

func (c *ControllerApp) DisableItem(ctx *gin.Context) {
	var dto dto.ItemDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success := c.ItemService.DisableItem(dto.ID, dto.Enable)
	if !success {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Disable Item"})
		return
	}
	ctx.JSON(http.StatusCreated, dto)

}

func (c *ControllerApp) FindItems(ctx *gin.Context) {
	var dto dto.ItemDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	itemsEntity := c.ItemService.FindItem(helpers.ToItemEntity(&dto))
	dtoList := helpers.ToItemDtoList(&itemsEntity)

	ctx.JSON(http.StatusCreated, dtoList)

}
