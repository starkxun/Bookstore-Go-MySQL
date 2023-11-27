package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){		//数据库这里看一下文档，可能存在问题
	d, err := gorm.Open("mysql","book:123456@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)		//返回让程序崩溃的错误
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}

