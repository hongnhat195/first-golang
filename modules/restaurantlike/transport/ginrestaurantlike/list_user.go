package ginrestaurantlike

import (
	"net/http"

	"github.com/hongnhat195/first-golang/common"
	rstlikebiz "github.com/hongnhat195/first-golang/modules/restaurantlike/biz"
	restaurantlikemodel "github.com/hongnhat195/first-golang/modules/restaurantlike/model"
	restaurantlikestorage "github.com/hongnhat195/first-golang/modules/restaurantlike/storage"

	"github.com/gin-gonic/gin"
	"github.com/hongnhat195/first-golang/component"
)

// GET /v1/restaurants/:id/liked-users

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))
		//var filter restaurantlikemodel.Filter
		//
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessReponse(result, paging, filter))
	}
}

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))
		//var filter restaurantlikemodel.Filter
		//
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			UserId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewListResraurantLikedByUser(store)

		result, err := biz.ListRestaurants(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		// for i := range result {
		// 	result[i].Mask(false)
		// }

		c.JSON(http.StatusOK, common.NewSuccessReponse(result, paging, filter))
	}
}
