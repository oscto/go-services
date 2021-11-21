package repository

import (
	"log"

	"account.oscto.icu/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func Register() *Repository {

	var repository = &Repository{
		db: connection(),
	}

	return repository
}

func connection() *gorm.DB {

	dsn := "host=192.168.123.10 user=postgres password=password dbname=account port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Panicln("数据库连接错误,", err)
	}
	initial(db)

	return db
}

func initial(db *gorm.DB) {

	if err := db.AutoMigrate(&model.Users{}); err != nil {
		log.Panicln("数据库表初始化错误")
	}
}
