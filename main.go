package main

import (
	"log"
	"os"

	"net/http"

	"github.com/hongnhat195/first-golang/component"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hongnhat195/first-golang/modules/restaurants/restauranttransport/ginrestaurant"
)

func main() {
	error := godotenv.Load(".env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION_STR")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PONG",
		})
	})
	appCtx := component.NewAppContext(db)
	restaurant := r.Group("/restaurants")
	{
		restaurant.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurant.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	}

	return r.Run()
}
