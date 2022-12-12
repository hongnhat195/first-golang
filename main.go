package main

import (
	"log"
	"os"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "go_delivery.restaurants"
}

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

	restaurant := r.Group("/restaurants")
	{
		restaurant.POST("", func(c *gin.Context) {
			var data Restaurant

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			if err := db.Create(&data).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})

		restaurant.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			var data Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, data)
		})

		restaurant.GET("", func(c *gin.Context) {
			var data []Restaurant
			type Filter struct {
				CityId int `json:"city_id" form: "city_id"`
			}

			var filter Filter
			c.ShouldBind(&filter)

			newDb := db
			if filter.CityId > 0 {
				newDb = db.Where("cituy_id = ?", filter.CityId)
			}

			if err := newDb.Find(&data).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})
	}

	return r.Run()
}
