package model

import (
	"time"
	"gorm.io/gorm"
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
	gorm.Model
	Name        string
	Category    string
	Author      string
	ImageLink   string
	Description string
}
