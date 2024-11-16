package typedefine

func ResponseHttpErrRsp(message string) HttpErrRsp {
	httpRsp := HttpErrRsp{}
	httpRsp.Code = 40020
	httpRsp.Message = message
	return httpRsp
}

func ResponseHttpSucRsp(data interface{}) HttpSucRsp {
	httpRsp := HttpSucRsp{}
	httpRsp.Code = 200
	httpRsp.Message = "ok"
	httpRsp.Data = data
	return httpRsp
}
