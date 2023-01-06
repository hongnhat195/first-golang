package restaurantmodel

import (
	"errors"
	"strings"

	"github.com/hongnhat195/first-golang/common"
)

const EntityName = "restaurants"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	// Id              int            `json:"id" gorm:"column:id;"`
	Name      string             `json:"name" gorm:"column:name;"`
	UserId    int                `json:"-" gorm:"column:owner_id"`
	Addr      string             `json:"address" gorm:"column:addr;"`
	Logo      *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover     *common.Images     `json:"cover" gorm:"column:cover;"`
	LikeCount int                `json:"like_count" gorm:"column:liked_count"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	UserId          int            `json:"-" gorm:"column:owner_id"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {

		return errors.New("Restaurant name can't be blank")
	}
	return nil
}

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}
