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
