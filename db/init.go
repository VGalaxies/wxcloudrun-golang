package db

import (
	"fmt"
	"os"
	"time"

	"wxcloudrun-golang/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB

// Init 初始化数据库
func Init() error {

	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "golang_demo"
	}
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	fmt.Println("start init mysql with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})
	if err != nil {
		fmt.Println("DB Open error,err=", err.Error())
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

  // init table
  if err = InitBook(); err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
  } 

  if err = InitCategory(); err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
  } 

	fmt.Println("finish init mysql with ", source)
	return nil
}

func InitBook() error {
	migrator := dbInstance.Migrator()
	if !migrator.HasTable(&model.BookModel{}) {
    err := migrator.CreateTable(model.BookModel{})
    if err != nil {
      return err
    }
	}

	return nil
}

func InitCategory() error {
	migrator := dbInstance.Migrator()
	if !migrator.HasTable(&model.CategoryModel{}) {
    err := migrator.CreateTable(model.CategoryModel{})
    if err != nil {
      return err
    }
	}

	return nil
}


// Get ...
func Get() *gorm.DB {
	return dbInstance
}
