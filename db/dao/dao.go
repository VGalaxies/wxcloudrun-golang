package dao

import (
	"errors"
	"fmt"
	"strconv"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

func (imp *BookInterfaceImp) GetBookByName(name string) (*model.BookModel, error) {
	var err error
	var book = new(model.BookModel)

	cli := db.Get()
	tx := cli.Model(&model.BookModel{}).Where("name = ?", name).First(book)
	if tx.RowsAffected == 0 {
		err = errors.New("book record not found")
		return nil, err
	}
	err = tx.Error

	return book, err
}

func (imp *BookInterfaceImp) GetBookByNameFzf(name string) (*[]model.BookModel, error) {
	var err error
	var books = new([]model.BookModel)

	cli := db.Get()
	tx := cli.Model(&model.BookModel{}).Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(books)

	if tx.RowsAffected == 0 {
		err = errors.New("book record not found")
		return nil, err
	}
	err = tx.Error

	return books, err
}

func (imp *BookInterfaceImp) GetBookByCategory(categoryIdStr string) (*[]model.BookModel, error) {
	var err error
	var books = new([]model.BookModel)

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		err = errors.New("invalid categoryId")
		return nil, err
	}

	cli := db.Get()
	tx := cli.Model(&model.BookModel{}).Where("category_id = ?", categoryId).Find(books)

	if tx.RowsAffected == 0 {
		err = errors.New("book record not found")
		return nil, err
	}
	err = tx.Error

	return books, err
}

// ----------------------- //

func (imp *CategoryInterfaceImp) GetCategory(categoryIdStr string) (*model.CategoryModel, error) {
	var err error
	var category = new(model.CategoryModel)

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		err = errors.New("invalid categoryId")
		return nil, err
	}

	cli := db.Get()
	tx := cli.Model(&model.CategoryModel{}).Where("id = ?", categoryId).First(category)
	if tx.RowsAffected == 0 {
		err = errors.New("category not found")
		return nil, err
	}
	err = tx.Error

	return category, err
}

func (imp *CategoryInterfaceImp) GetCategoryAll() (*[]model.CategoryModel, error) {
	var err error
	var categories = new([]model.CategoryModel)

	cli := db.Get()
	tx := cli.Model(&model.CategoryModel{}).Find(categories)
	if tx.RowsAffected == 0 {
		err = errors.New("no category yet")
		return nil, err
	}
	err = tx.Error

	return categories, err
}

// ----------------------- //

func (imp *UserInterfaceImp) SetUserInfo(openid string, nickname string, avatar string) error {
	var err error
	var user = new(model.UserModel)

	cli := db.Get()
	// ignore return value
	cli.Model(&model.UserModel{}).Where("open_id = ?", openid).Delete(user)
	tx := cli.Model(&model.UserModel{}).Create(map[string]interface{}{
		"open_id":    openid,
		"nick_name":  nickname,
		"avatar_url": avatar,
	})
	err = tx.Error

	return err
}

func (imp *UserInterfaceImp) GetUserInfo(openid string) (*model.UserModel, error) {
	var err error
	var user = new(model.UserModel)

	cli := db.Get()
	tx := cli.Model(&model.UserModel{}).Where("open_id = ?", openid).First(user)
	if tx.RowsAffected == 0 {
		err = errors.New("user not found")
		return nil, err
	}
	err = tx.Error

	return user, err
}

// ----------------------- //

func (imp *CommentInterfaceImp) SetCommentInfo(userId string, bookIdStr string, comment string) error {
	var err error

	bookId, err := strconv.ParseInt(bookIdStr, 10, 32)
	if err != nil {
		err = errors.New("invalid bookId")
		return err
	}

	cli := db.Get()
	// TODO - integrity checking
	tx := cli.Model(&model.CommentModel{}).Create(map[string]interface{}{
		"user_id": userId,
		"book_id": bookId,
		"comment": comment,
	})
	err = tx.Error

	return err
}

func (imp *CommentInterfaceImp) GetCommentInfoByUser(userId string) (*[]model.CommentModel, error) {
	var err error
	var comments = new([]model.CommentModel)

	cli := db.Get()
	tx := cli.Model(&model.CommentModel{}).Where("user_id = ?", userId).Find(comments)

	if tx.RowsAffected == 0 {
		err = errors.New("comment record not found")
		return nil, err
	}
	err = tx.Error

	return comments, err
}

func (imp *CommentInterfaceImp) GetCommentInfoByBook(bookIdStr string) (*[]model.CommentModel, error) {
	var err error
	var comments = new([]model.CommentModel)

	bookId, err := strconv.ParseInt(bookIdStr, 10, 32)
	if err != nil {
		err = errors.New("invalid bookId")
		return nil, err
	}

	cli := db.Get()
	tx := cli.Model(&model.CommentModel{}).Where("book_id = ?", bookId).Find(comments)

	if tx.RowsAffected == 0 {
		err = errors.New("comment record not found")
		return nil, err
	}
	err = tx.Error

	return comments, err
}
