package dao

import (
	"wxcloudrun-golang/db/model"
)

// 书籍数据模型接口
type BookInterface interface {
	GetBookByName(string) (*model.BookModel, error)
	GetBookByNameFzf(string) (*[]model.BookModel, error)
	GetBookByCategory(string) (*[]model.BookModel, error)
}

// 书籍数据模型实现
type BookInterfaceImp struct{}

// 书籍数据模型实现实例
var BookImp BookInterface = &BookInterfaceImp{}

// ----------------------- //

// 书籍分类数据模型接口
type CategoryInterface interface {
	GetCategory(string) (*model.CategoryModel, error)
	GetCategoryAll() (*[]model.CategoryModel, error)
}

// 书籍分类数据模型实现
type CategoryInterfaceImp struct{}

// 书籍分类数据模型实现实例
var CategoryImp CategoryInterface = &CategoryInterfaceImp{}

// ----------------------- //

// 用户数据模型接口
type UserInterface interface {
	SetUserInfo(string, string, string) error
	GetUserInfo(string) (*model.UserModel, error)
}

type UserInterfaceImp struct{}

var UserImp UserInterface = &UserInterfaceImp{}
