package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hongnhat195/first-golang/common"
	"github.com/hongnhat195/first-golang/component"
	"github.com/hongnhat195/first-golang/component/hasher"
	"github.com/hongnhat195/first-golang/modules/user/userbiz"
	"github.com/hongnhat195/first-golang/modules/user/usermodel"
	"github.com/hongnhat195/first-golang/modules/user/userstorage"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
