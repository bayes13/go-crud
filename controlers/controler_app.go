package controlers

import (
	"go-crud/helpers"
	"go-crud/models/dto"
	"go-crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerApp struct {
	Service services.ItemService
}

func NewControllerApp(service services.ItemService) *ControllerApp {
	return &ControllerApp{Service: service}
}

func (c *ControllerApp) CreateItem(ctx *gin.Context) {
	var dto dto.ItemDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success := c.Service.CreateItem(helpers.ToItemEntity(&dto))
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
	success := c.Service.UpdateItem(helpers.ToItemEntity(&dto))
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
	success := c.Service.DisableItem(dto.ID, dto.Enable)
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
	itemsEntity := c.Service.FindItem(helpers.ToItemEntity(&dto))
	dtoList := helpers.ToItemDtoList(&itemsEntity)

	ctx.JSON(http.StatusCreated, dtoList)

}
