package main

import (
	"Restorant-management/database"
	"Restorant-management/middleware"
	"Restorant-management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4/database"
	"go.mongodb.org/mongo-driver/mongo"
)
var foodCollection *mongo.Collection = database.OpenCollection(database.client,"food")
func main() {
	port :=os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())


	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}