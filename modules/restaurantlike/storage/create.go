package restaurantlikestorage

import (
	"context"

	restaurantlikemodel "github.com/hongnhat195/first-golang/modules/restaurantlike/model"

	"github.com/hongnhat195/first-golang/common"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
