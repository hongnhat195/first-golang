package main

import (
	"log"
	"os"

	"github.com/hongnhat195/first-golang/component"
	"github.com/hongnhat195/first-golang/component/uploadprovider"
	"github.com/hongnhat195/first-golang/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hongnhat195/first-golang/modules/restaurantlike/transport/ginrestaurantlike"
	"github.com/hongnhat195/first-golang/modules/restaurants/restauranttransport/ginrestaurant"
	"github.com/hongnhat195/first-golang/modules/upload/uploadtransport/ginupload"
	"github.com/hongnhat195/first-golang/modules/user/usertransport/ginuser"
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

	secretKey := os.Getenv("JWT_SECRET_KEY")

	s3provider := uploadprovider.NewS3Provider(s3_bucket, s3_region, s3_api_key, s3_secret_key, s3_domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := runService(db, s3provider, secretKey); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB, upprovider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppContext(db, upprovider, secretKey)
	r := gin.Default()

	r.Use(middlewares.Recover(appCtx))

	//CRUD

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))

	v1.POST("/login", ginuser.Login(appCtx))

	v1.GET("/profile", middlewares.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	v1.POST("/upload", ginupload.Upload(appCtx))

	restaurant := v1.Group("/restaurants")
	{
		restaurant.POST("", middlewares.RequiredAuth(appCtx), ginrestaurant.CreateRestaurant(appCtx))

		restaurant.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurant.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurant.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurant.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))

		restaurant.GET("/:id/liked-restaurant", ginrestaurantlike.ListRestaurant(appCtx))

		restaurant.POST("/:id/like", middlewares.RequiredAuth(appCtx), ginrestaurantlike.UserLikeRestaurant(appCtx))

		restaurant.DELETE("/:id/unlike", middlewares.RequiredAuth(appCtx), ginrestaurantlike.UserUnlikeRestaurant(appCtx))

	}

	return r.Run()
}
