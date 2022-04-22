package model

import (
	"time"
	"wxcloudrun-golang/db"
)

// CounterModel 计数器模型
type CounterModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Count     int32     `gorm:"column:count" json:"count"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// demo
type BookModel struct {
	Id          int32
	Name        string
	Category    int32
	Author      string
	Description string
}

type CategoryModel struct {
	Id   int32
	Name string
}

func InitBook() error {
  res := db.Get().Create(&[]BookModel{
    {Id: 1, Name: "Operating Systems: Three Easy Pieces", Category: 1, Author: "Remzi H. Arpaci-Dusseau", Description: ""},
    {Id: 2, Name: "Models of Computation", Category: 2, Author: "Jeff Erickson", Description: ""},
    {Id: 3, Name: "Fundamentals of Computer Graphics", Category: 3, Author: "Peter Shirley", Description: ""},
	})

  return res.Error;
}

func InitCategory() error {
  res := db.Get().Create(&[]CategoryModel{
    {Id: 1, Name: "Operating System"},
    {Id: 2, Name: "Theoretical Computer Science"},
    {Id: 3, Name: "Computer Graphics"},
	})

	return res.Error
}

