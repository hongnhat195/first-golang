package restaurantstore

import (
	"context"

	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db
	if err := db.Table(restaurantmodel.RestaurantUpdate{}.TableName()).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
