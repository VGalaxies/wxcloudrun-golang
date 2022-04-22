package dao

import (
	"errors"
	"fmt"
	"strconv"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

func (imp *BookInterfaceImp) GetBookByName(name string) (*[]model.BookModel, error) {
	var err error
	var books = new([]model.BookModel)

	cli := db.Get()
	tx := cli.Model(&model.BookModel{}).Where("name = ?", name).Find(books)
	if tx.RowsAffected != 1 {
		err = errors.New("duplicated book name")
		return nil, err
	}
	err = tx.Error

	return books, err
}

func (imp *BookInterfaceImp) GetBookByNameFzf(name string) (*[]model.BookModel, error) {
	var err error
	var books = new([]model.BookModel)

	cli := db.Get()
	err = cli.Model(&model.BookModel{}).Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(books).Error

	return books, err
}

func (imp *BookInterfaceImp) GetBookByNameCate(categoryIdStr string) (*[]model.BookModel, error) {
	var err error
	var books = new([]model.BookModel)

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		err = errors.New("invalid categoryId")
		return nil, err
	}

	cli := db.Get()
	err = cli.Model(&model.BookModel{}).Where("Category = ?", categoryId).Find(books).Error

	return books, err
}
