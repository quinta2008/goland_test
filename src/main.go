package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"xorm.io/xorm"
)

type test1 struct {
	age    int
	name   string
	height int32
}

func dbInit() {
	engine, _ := xorm.NewEngine("sqlite3", "./test.db")
	var a test1
	err := engine.Sync2(a)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func main() {
	fmt.Println("Hello, World!")
	dbInit()
	for true {
		time.Sleep(100)
	}
}
