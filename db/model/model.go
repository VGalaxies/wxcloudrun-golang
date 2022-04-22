package model

// 书籍模型
type BookModel struct {
	Id          int32
	Name        string
	Category    int32
	Author      string
	Description string
}

// 书籍分类模型
type CategoryModel struct {
	Id   int32
	Name string
}
