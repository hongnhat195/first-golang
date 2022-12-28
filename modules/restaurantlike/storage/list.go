package restaurantlikestorage

import "context"

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	for i, item := range ids {
		result[i] = item
	}

	return result, nil
}
