package customerrequest

import (
	"fmt"
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

// 定义带自定义参数的回调处理函数
func handleGetConfigValues(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var configValues typedefine.ConfigValues
		if err := c.ShouldBind(&configValues); err != nil {
			c.JSON(http.StatusBadRequest, typedefine.ResponseHttpErrRsp("Bind Json Failed"))
			return
		}
		// 调用函数读取数据
		configValuesRspList := []typedefine.ConfigValuesRsp{}
		err := typedefine.ReadJSONFileToStruct("conf/config_value.json", &configValuesRspList)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// 处理业务逻辑
		c.JSON(http.StatusOK, typedefine.ResponseHttpSucRsp(configValuesRspList))
	}
}

// 定义带自定义参数的回调处理函数
func handleGetBannerList(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bannerList typedefine.BannerList
		if err := c.ShouldBind(&bannerList); err != nil {
			c.JSON(http.StatusBadRequest, typedefine.ResponseHttpErrRsp("Bind Json Failed"))
			return
		}
		// 调用函数读取数据
		bannerListRsp := []typedefine.BannerListRsp{}
		err := typedefine.ReadJSONFileToStruct("conf/banner_list.json", &bannerListRsp)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// 处理业务逻辑
		c.JSON(http.StatusOK, typedefine.ResponseHttpSucRsp(bannerListRsp))
	}
}

// 定义带自定义参数的回调处理函数
func handleGetShopGoodList(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shopGoodsList typedefine.ShopGoodsList
		if err := c.ShouldBind(&shopGoodsList); err != nil {
			c.JSON(http.StatusBadRequest, typedefine.ResponseHttpErrRsp("Bind Json Failed"))
			return
		}
		// 调用函数读取数据
		shopGoodsRspList := []typedefine.ShopGoodsListRsp{}
		err := typedefine.ReadJSONFileToStruct("conf/shop_goods_list.json", &shopGoodsRspList)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// 处理业务逻辑
		c.JSON(http.StatusOK, typedefine.ResponseHttpSucRsp(shopGoodsRspList))
	}
}

func HandleCustomerRequest(r *gin.Engine) {
	var dbEngine = dao.Engine.GetEngine()
	r.GET("/users/cms/category/info", handleGetCmsCategoryInfo(dbEngine))
	r.GET("/config/values", handleGetConfigValues(dbEngine))
	r.GET("/banner/list", handleGetBannerList(dbEngine))
	r.GET("/shop/goods/list/v2", handleGetShopGoodList(dbEngine))
}
