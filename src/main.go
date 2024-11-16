package main

import (
	"github.com/gin-gonic/gin"
	"golang_test/src/customerrequest"
	"golang_test/src/dao"
	"log"
)

func initHttpEngine() error {
	r := gin.Default()
	customerrequest.HandleCustomerRequest(r)
	// 注册回调处理函数
	// 启动HTTP服务器
	err := r.Run(":8080") // 监听并在 0.0.0.0:8080 启动服务
	return err
}

func main() {
	dao.Engine.InitDataBase()
	err := initHttpEngine()
	if err != nil {
		log.Fatal(err)
	}
}
