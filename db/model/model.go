package model

// 书籍模型
type BookModel struct {
	Id          int32
	Name        string `gorm:"unique"`
	CategoryId  int32
	Author      string
	Description string
	ImageUrl    string
	Comments    []CommentModel
}

// 书籍分类模型
type CategoryModel struct {
	Id   int32
	Name string
}

// 用户模型
type UserModel struct {
	Id        int32
	OpenId    string
	NickName  string
	AvatarUrl string
	// RecentBooks   []BookModel
	// FavoriteBooks []BookModel
}

// 评论模型
type CommentModel struct {
	Id          int32
	UserModelId int32
	BookModelId int32
	Comment     string
}
