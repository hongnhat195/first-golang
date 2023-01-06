package ginrestaurantlike

import (
	"net/http"

	rstlikebiz "github.com/hongnhat195/first-golang/modules/restaurantlike/biz"
	restaurantlikemodel "github.com/hongnhat195/first-golang/modules/restaurantlike/model"
	restaurantlikestorage "github.com/hongnhat195/first-golang/modules/restaurantlike/storage"

	"github.com/gin-gonic/gin"
	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/component"
)

// POST /v1/restaurants/:id/like

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
