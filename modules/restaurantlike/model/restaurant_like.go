package restaurantlikemodel

import (
	"fmt"
	"time"

	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int        `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int        `json:"user_id" gorm:"column:user_id"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
	// UpdatedAt    *time.Time                  `json:"updated_at" gorm:"column:updated_at"`
	User       *common.SimpleUser          `json:"user" gorm:"preload:false;"`
	Restaurant *restaurantmodel.Restaurant `json:"restaurant" gorm:"preload:false;"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this restaurant"),
		fmt.Sprintf("ErrCannotUnlikeRestaurant"),
	)
}
