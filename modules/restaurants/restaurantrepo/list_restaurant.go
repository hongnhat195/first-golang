package restaurantrepo

import (
	"context"

	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListDataByCondition(ctx, nil, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	// ids := make([]int, len(result))

	// for i := range result {
	// 	ids[i] = result[i].Id
	// }

	// mapResLike, err := repo.likeStore.GetRestaurantLikes(ctx, ids)

	// if err != nil {
	// 	log.Println("Cannot get restaurant likes:", err)
	// }

	// if v := mapResLike; v != nil {
	// 	for i, item := range result {
	// 		result[i].LikeCount = mapResLike[item.Id]
	// 	}
	// }

	return result, nil
}
