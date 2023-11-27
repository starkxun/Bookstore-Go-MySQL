package models

import (
	"github.com/jinzhu/gorm"
	"github.com/starkxun/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	//gorm.Model是gorm库提供的结构体，用于嵌套在其他结构体中
	//提供了一下常见的数据库字段，如ID、CreatedAt、UpdatedAt和DeletedAt，用于支持数据库操作中的常见需求，如主键、创建时间、更新时间和软删除
	gorm.Model
	Name string `gorm:"":json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})	//根据Book结构自动建表
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)		//给结构体对象b标记一个新的记录，通常用于设置主键信息
	db.Create(&b)		//将结构体对象b插入到数据库中
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book,*gorm.DB){
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db 
} 

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Delete(book)
	return book
}