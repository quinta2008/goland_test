package customerrequest

import (
	"github.com/gin-gonic/gin"
	"golang_test/src/dao"
	"golang_test/src/typedefine"
	"net/http"
	"xorm.io/xorm"
)

// 定义带自定义参数的回调处理函数
func handleGetCmsCategoryInfo(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmsCategoryInfo typedefine.CmsCategoryInfo
		if err := c.ShouldBind(&cmsCategoryInfo); err != nil {
			c.JSON(http.StatusBadRequest, typedefine.ResponseHttpErrRsp("Bind Json Failed"))
			return
		}
		// 处理业务逻辑
		c.JSON(http.StatusOK, typedefine.ResponseHttpSucRsp(cmsCategoryInfo))
	}
}

func HandleCustomerRequest(r *gin.Engine) {
	var dbEngine = dao.Engine.GetEngine()
	r.GET("/users/cms/category/info", handleGetCmsCategoryInfo(dbEngine))
}
