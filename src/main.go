package main

import (
	"fmt"
	_ "github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	for true {
		time.Sleep(100)
	}
}
