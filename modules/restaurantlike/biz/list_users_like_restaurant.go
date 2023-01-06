package rstlikebiz

import (
	"context"

	restaurantlikemodel "github.com/hongnhat195/first-golang/modules/restaurantlike/model"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"

	"github.com/hongnhat195/first-golang/common"
)

type ListUserLikeRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}

type ListRestaurantLikedByUser interface {
	GetRestaurantLikedByUser(ctx context.Context, conditions map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantLikedByUserBiz struct {
	store ListRestaurantLikedByUser
}

func NewListResraurantLikedByUser(store ListRestaurantLikedByUser) *listRestaurantLikedByUserBiz {
	return &listRestaurantLikedByUserBiz{store: store}
}

func (biz *listUserLikeRestaurantBiz) ListUsers(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}

func (biz *listRestaurantLikedByUserBiz) ListRestaurants(ctx context.Context, filter *restaurantlikemodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	restaurant, err := biz.store.GetRestaurantLikedByUser(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return restaurant, nil
}
