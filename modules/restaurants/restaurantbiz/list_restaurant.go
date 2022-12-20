package restaurantbiz

import (
	"context"

	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewlistRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {

	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return result, err
}
