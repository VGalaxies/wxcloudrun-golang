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

// 用户数据模型实现
type UserInterfaceImp struct{}

// 用户数据模型实现实例
var UserImp UserInterface = &UserInterfaceImp{}

// ----------------------- //

// 评论数据模型接口
type CommentInterface interface {
	SetCommentInfo(string, string, string) error
	GetCommentInfoByUser(string) (*[]model.CommentModel, error)
	GetCommentInfoByBook(string) (*[]model.CommentModel, error)
}

// 评论数据模型实现
type CommentInterfaceImp struct{}

// 评论数据模型实现实例
var CommentImp CommentInterface = &CommentInterfaceImp{}

// ----------------------- //

// 收藏数据模型接口
type CollectionInterface interface {
	SetCollectionInfo(string, string) error
	GetCollectionInfoByUser(string) (*[]model.CollectionModel, error)
	GetCollectionInfoByBook(string) (*[]model.CollectionModel, error)
}

// 收藏数据模型实现
type CollectionInterfaceImp struct{}

// 收藏数据模型实现实例
var CollectionImp CollectionInterface = &CollectionInterfaceImp{}
