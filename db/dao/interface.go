package dao

import (
	"wxcloudrun-golang/db/model"
)

// 书籍数据模型接口
type BookInterface interface {
	GetBookByName(string) (*[]model.BookModel, error)
	GetBookByNameFzf(string) (*[]model.BookModel, error)
	GetBookByNameCate(string) (*[]model.BookModel, error)
}

// 书籍数据模型实现
type BookInterfaceImp struct{}

// 书籍数据模型实现实例
var BookImp BookInterface = &BookInterfaceImp{}
