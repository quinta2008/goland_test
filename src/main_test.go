package main

import (
	"fmt"
	"testing"
	"xorm.io/xorm"
)

type test1 struct {
	Age    int
	Name   string
	Height int32
}

func dbInit() {
	engine, _ := xorm.NewEngine("sqlite3", "./test.db")
	defer engine.Close()
	var a test1
	a.Age = 10
	a.Name = "qiaotianshu"
	err := engine.Sync2(a)
	if err != nil {
		fmt.Println(err)
	}
	_, err = engine.Insert(a)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func dbSelect() {
	engine, _ := xorm.NewEngine("sqlite3", "./test.db")
	defer engine.Close()
	type test2 struct {
		Age int
	}
	var b []test2
	err := engine.Table("test1").Find(&b)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func TestDbInsert(t *testing.T) {
	dbInit()
	dbSelect()
}
