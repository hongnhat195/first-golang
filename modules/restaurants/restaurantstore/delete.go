package restaurantstore

import (
	"context"

	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"
)

func (s *sqlStore) SoftDeleteData(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
