package model

import (
	"wxcloudrun-golang/db"
)

func InitBook() error {
	migrator := db.Get().Migrator()
	if !migrator.HasTable(&BookModel{}) {
    res := migrator.CreateTable(BookModel{})
    if res != nil {
      return res
    }
	}

	res := db.Get().Create(&[]BookModel{
		{Id: 1, Name: "Operating Systems: Three Easy Pieces", Category: 1, Author: "Remzi H. Arpaci-Dusseau", Description: ""},
		{Id: 2, Name: "Models of Computation", Category: 2, Author: "Jeff Erickson", Description: ""},
		{Id: 3, Name: "Fundamentals of Computer Graphics", Category: 3, Author: "Peter Shirley", Description: ""},
	})

	return res.Error
}

func InitCategory() error {
	migrator := db.Get().Migrator()
	if !migrator.HasTable(&CategoryModel{}) {
    res := migrator.CreateTable(CategoryModel{})
    if res != nil {
      return res
    }
	}

	res := db.Get().Create(&[]CategoryModel{
		{Id: 1, Name: "Operating System"},
		{Id: 2, Name: "Theoretical Computer Science"},
		{Id: 3, Name: "Computer Graphics"},
	})

	return res.Error
}

