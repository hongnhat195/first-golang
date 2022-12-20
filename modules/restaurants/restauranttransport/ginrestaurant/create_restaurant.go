package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/component"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantbiz"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"

	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantstore"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
