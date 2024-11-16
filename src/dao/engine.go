package dao

import "xorm.io/xorm"

type DbEngine struct {
	engine *xorm.Engine
}

var Engine DbEngine

func (c *DbEngine) InitDataBase() *xorm.Engine {
	c.engine, _ = xorm.NewEngine("sqlite3", "./test.db")
	return c.engine
}
func (c *DbEngine) GetEngine() *xorm.Engine {
	return c.engine
}
