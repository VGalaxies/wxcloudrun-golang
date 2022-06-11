package model

import "gorm.io/gorm"

// 书籍模型
type BookModel struct {
	Id          int32
	Name        string `gorm:"unique"`
	CategoryId  int32
	Author      string
	Description string
	ImageUrl    string
}

// 书籍分类模型
type CategoryModel struct {
	Id   int32
	Name string
}

// 用户模型
type UserModel struct {
	OpenId    string `gorm:"primaryKey"`
	NickName  string
	AvatarUrl string
}

// 评论模型
type CommentModel struct {
	gorm.Model
	UserId  string
	BookId  int32
	Comment string
}

// 收藏模型
type CollectionModel struct {
	gorm.Model
	UserId string
	BookId int32
}
