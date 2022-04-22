package model

// 书籍模型
type BookModel struct {
	Id          int32
	Name        string
	CategoryId  int32
	Author      string
	Description string
	ImageLink   string
}

// 书籍分类模型
type CategoryModel struct {
	Id   int32
	Name string
}
