package restaurantbiz

import (
	"context"
	"errors"

	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	SoftDeleteData(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewdeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldata, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if oldata.Status == 0 {
		return errors.New("data deleted")
	}
	if er := biz.store.SoftDeleteData(ctx, id); err != nil {
		return er
	}

	return nil
}
