package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/component"
	restaurantlikestorage "github.com/hongnhat195/first-golang/modules/restaurantlike/storage"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantbiz"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantrepo"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantstore"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)

		biz := restaurantbiz.NewlistRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(true)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessReponse(result, paging, filter))
	}
}
