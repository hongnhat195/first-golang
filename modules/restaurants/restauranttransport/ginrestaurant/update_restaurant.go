package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/component"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantbiz"
	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantmodel"

	"github.com/hongnhat195/first-golang/modules/restaurants/restaurantstore"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
