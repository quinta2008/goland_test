package typedefine

type HttpErrRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HttpSucRsp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CmsCategoryInfo struct {
	Id              string `json:"id"`
	YuyueStatistics string `json:"yuyuestatistics"`
	YuyueUid        string `json:"yuyueuid"`
	Token           string `json:"token"`
}

// 站点设置
type ConfigValues struct {
	Keys string `json:"keys"`
}

type ConfigValuesRsp struct {
	DateUpdate string `json:"dateUpdate,omitempty"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	Remark     string `json:"remark"`
}

type BannerList struct {
	Type string `json:"type"` // 使用 JSON 标签映射 "type" 字段
}

type BannerListRsp struct {
	BusinessId int    `json:"businessId"`
	DateAdd    string `json:"dateAdd"`
	DateUpdate string `json:"dateUpdate"`
	Id         int    `json:"id"`
	LinkType   int    `json:"linkType"`
	LinkUrl    string `json:"linkUrl"`
	Paixu      int    `json:"paixu"`
	PicUrl     string `json:"picUrl"`
	ShopId     int    `json:"shopId"`
	Status     int    `json:"status"`
	StatusStr  string `json:"statusStr"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	UserId     int    `json:"userId"`
}

// 定义请求结构体
type ShopGoodsList struct {
	RecommendStatus int `json:"recommendStatus"`
}

// 定义响应结构体
type ShopGoodsListRsp struct {
	AfterSale            string `json:"afterSale"`
	BarCode              string `json:"barCode,omitempty"`
	CategoryId           int    `json:"categoryId"`
	Characteristic       string `json:"characteristic"`
	Commission           int    `json:"commission"`
	CommissionSettleType int    `json:"commissionSettleType"`
	CommissionType       int    `json:"commissionType"`
	CommissionUserType   int    `json:"commissionUserType"`
	DateAdd              string `json:"dateAdd"`
	DateUpdate           string `json:"dateUpdate"`
	DateStart            string `json:"dateStart,omitempty"`
	DiscountPrice        int    `json:"discountPrice"`
	FxType               int    `json:"fxType"`
	GotScore             int    `json:"gotScore"`
	GotScoreType         int    `json:"gotScoreType"`
	HasAddition          bool   `json:"hasAddition"`
	HasTourJourney       bool   `json:"hasTourJourney"`
	Hidden               int    `json:"hidden"`
	Id                   int    `json:"id"`
	IotControl           bool   `json:"iotControl"`
	Kanjia               bool   `json:"kanjia"`
	KanjiaPrice          int    `json:"kanjiaPrice"`
	Limitation           bool   `json:"limitation"`
	LogisticsId          int    `json:"logisticsId"`
	MaxCoupons           int    `json:"maxCoupons"`
	Miaosha              bool   `json:"miaosha"`
	MinBuyNumber         int    `json:"minBuyNumber"`
	MinPrice             int    `json:"minPrice"`
	MinScore             int    `json:"minScore"`
	Name                 string `json:"name"`
	NumberFav            int    `json:"numberFav"`
	NumberGoodReputation int    `json:"numberGoodReputation"`
	NumberOrders         int    `json:"numberOrders"`
	NumberReputation     int    `json:"numberReputation"`
	NumberSells          int    `json:"numberSells"`
	OriginalPrice        int    `json:"originalPrice"`
	Overseas             bool   `json:"overseas"`
	Paixu                int    `json:"paixu"`
	Persion              int    `json:"persion"`
	Pic                  string `json:"pic"`
	Pingtuan             bool   `json:"pingtuan"`
	PingtuanPrice        int    `json:"pingtuanPrice"`
	PriceShopSell        int    `json:"priceShopSell"`
	RecommendStatus      int    `json:"recommendStatus"`
	RecommendStatusStr   string `json:"recommendStatusStr"`
	SeckillBuyNumber     int    `json:"seckillBuyNumber"`
	SellEnd              bool   `json:"sellEnd"`
	SellStart            bool   `json:"sellStart"`
	ShopId               int    `json:"shopId"`
	Status               int    `json:"status"`
	StatusStr            string `json:"statusStr"`
	StoreAlert           bool   `json:"storeAlert"`
	StoreAlertNum        int    `json:"storeAlertNum"`
	Stores               int    `json:"stores"`
	Stores0Unsale        bool   `json:"stores0Unsale"`
	StoresExt1           int    `json:"storesExt1"`
	StoresExt2           int    `json:"storesExt2"`
	SubName              string `json:"subName,omitempty"`
	Tax                  int    `json:"tax"`
	Type                 int    `json:"type"`
	Unit                 string `json:"unit"`
	UnusefulNumber       int    `json:"unusefulNumber"`
	UsefulNumber         int    `json:"usefulNumber"`
	UserId               int    `json:"userId"`
	VetStatus            int    `json:"vetStatus"`
	Views                int    `json:"views"`
	Weight               int    `json:"weight"`
}
