package main

import (
	"go-crud/config"
	"go-crud/controlers"
	"go-crud/repositories"
	"go-crud/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectionDataBase()
	config.InitDatabase()

	router := gin.Default()

	itemServ := services.NewItemService(repositories.NewItemRepository())
	controller := controlers.NewControllerApp(itemServ)

	controlers.SetUpRoutes(router, controller)

	log.Println("Starting Server on port 8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server fauiled: %v", err)
	}

}
