package controlers

import (
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine, c *ControllerApp) {
	r.POST("/item/create", c.CreateItem)
	r.POST("/item/update", c.UpdateItem)
	r.POST("/item/disable", c.UpdateItem)
	r.POST("/item/find", c.FindItems)

}
