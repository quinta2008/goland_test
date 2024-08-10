package main

import (
	"encoding/json"
	_ "github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

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

// getUserHandler 处理 GET 请求，返回用户列表
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 将用户列表编码为 JSON 并写入响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// createUserHandler 处理 POST 请求，解析 JSON 请求体并创建新用户
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体中的 JSON 数据
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 模拟创建用户，将其添加到内存中的用户数据中
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	// 返回创建的用户信息
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func loadRouter() {
	// 路由设置
	http.HandleFunc("/users", getUserHandler)
	http.HandleFunc("/users/create", createUserHandler)
}

func initHttpEngine() error {
	// 服务器端口
	port := ":8080"

	// 启动 HTTP 服务器
	log.Printf("Starting server on http://localhost%v\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
	return err
}

func main() {
	loadRouter()
	err := initHttpEngine()
	if err != nil {
		log.Fatal(err)
	}
}
