package main

import (
	_ "github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"xorm.io/xorm"
)

var engine *xorm.Engine

// User 定义一个用户结构体用于 JSON 数据的解析和响应
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 内存中的用户数据
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// 定义带自定义参数的回调处理函数
func createUserHandler(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, ResponseHttpErrRsp("Bind Json Failed"))
			return
		}
		// 处理业务逻辑
		c.JSON(http.StatusOK, ResponseHttpSucRsp(user))
	}
}

func getUserHandler(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.BindQuery(&user); err != nil {
			c.JSON(http.StatusBadRequest, ResponseHttpErrRsp("Bind Json Failed"))
			return
		}
	}
}

func initHttpEngine() error {
	r := gin.Default()

	// 注册回调处理函数
	r.POST("/users/create", createUserHandler(engine))
	r.GET("/users/query", getUserHandler(engine))
	// 启动HTTP服务器
	err := r.Run(":8080") // 监听并在 0.0.0.0:8080 启动服务
	return err
}

func initDataBase() {
	engine, _ = xorm.NewEngine("sqlite3", "./test.db")
	return
}

func main() {
	initDataBase()
	err := initHttpEngine()
	if err != nil {
		log.Fatal(err)
	}
}
