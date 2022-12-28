package main

import (
	"log"
	"os"

	"net/http"

	"github.com/hongnhat195/first-golang/component"
	"github.com/hongnhat195/first-golang/component/uploadprovider"
	"github.com/hongnhat195/first-golang/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hongnhat195/first-golang/modules/restaurants/restauranttransport/ginrestaurant"
	"github.com/hongnhat195/first-golang/modules/upload/uploadtransport/ginupload"
)

func main() {
	error := godotenv.Load(".env")
	if error != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION_STR")

	s3_bucket := os.Getenv("S3_BUCKET_NAME")
	s3_region := os.Getenv("S3_REGION")
	s3_api_key := os.Getenv("S3_ACCESS_KEY")
	s3_secret_key := os.Getenv("S3_SECRET_KEY")
	s3_domain := os.Getenv("S3_DOMAIN")

	s3provider := uploadprovider.NewS3Provider(s3_bucket, s3_region, s3_api_key, s3_secret_key, s3_domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3provider); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB, upprovider uploadprovider.UploadProvider) error {

	appCtx := component.NewAppContext(db, upprovider)
	r := gin.Default()
	r.Use(middlewares.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PONG",
		})
	})

	r.POST("/upload", ginupload.Upload(appCtx))

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
